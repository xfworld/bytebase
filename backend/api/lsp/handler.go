package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	lsp "github.com/bytebase/lsp-protocol"
	"github.com/pkg/errors"
	"github.com/sourcegraph/jsonrpc2"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/common/log"
	"github.com/bytebase/bytebase/backend/component/config"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	"github.com/bytebase/bytebase/backend/store"
)

type Method string

const (
	LSPMethodPing           Method = "$ping"
	LSPMethodInitialize     Method = "initialize"
	LSPMethodInitialized    Method = "initialized"
	LSPMethodShutdown       Method = "shutdown"
	LSPMethodExit           Method = "exit"
	LSPMethodCancelRequest  Method = "$/cancelRequest"
	LSPMethodSetTrace       Method = "$/setTrace"
	LSPMethodExecuteCommand Method = "workspace/executeCommand"
	LSPMethodCompletion     Method = "textDocument/completion"

	LSPMethodTextDocumentDidOpen   Method = "textDocument/didOpen"
	LSPMethodTextDocumentDidChange Method = "textDocument/didChange"
	LSPMethodTextDocumentDidClose  Method = "textDocument/didClose"
	LSPMethodTextDocumentDidSave   Method = "textDocument/didSave"

	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_publishDiagnostics.
	LSPMethodPublishDiagnostics Method = "textDocument/publishDiagnostics"

	// Custom Methods.
	// See dollar request: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#dollarRequests.
	LSPCustomMethodSQLStatementRanges Method = "$/textDocument/statementRanges"
)

// NewHandler creates a new Language Server Protocol handler.
func NewHandler(s *store.Store, profile *config.Profile) jsonrpc2.Handler {
	return lspHandler{Handler: jsonrpc2.HandlerWithError((&Handler{store: s, profile: profile}).handle)}
}

type lspHandler struct {
	jsonrpc2.Handler
}

func (h lspHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	if isFileSystemRequest(req.Method) {
		h.Handler.Handle(ctx, conn, req)
		return
	}
	go h.Handler.Handle(ctx, conn, req)
}

// Handler handles Language Server Protocol requests.
type Handler struct {
	mu       sync.Mutex
	fs       *MemFS
	init     *lsp.InitializeParams // set by LSPMethodInitialize request
	metadata *SetMetadataCommandArguments
	store    *store.Store

	shutDown bool
	profile  *config.Profile
	cancelF  sync.Map // map[jsonrpc2.ID]context.CancelFunc
}

// ShutDown shuts down the handler.
func (h *Handler) ShutDown() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.shutDown {
		slog.Warn("server received a shutdown request after it was already shut down.")
	}
	h.shutDown = true
	h.fs = nil
}

func (h *Handler) setMetadata(arg SetMetadataCommandArguments) {
	h.mu.Lock()
	defer h.mu.Unlock()
	tmp := arg
	h.metadata = &tmp
}

func (h *Handler) getDefaultDatabase() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.metadata == nil {
		return ""
	}
	return h.metadata.DatabaseName
}

func (h *Handler) getDefaultSchema() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.metadata == nil {
		return ""
	}
	return h.metadata.Schema
}

func (h *Handler) getInstanceID() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.metadata == nil {
		return ""
	}
	id, err := common.GetInstanceID(h.metadata.InstanceID)
	if err != nil {
		return ""
	}
	return id
}

func (h *Handler) getScene() base.SceneType {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.metadata == nil {
		return base.SceneTypeAll
	}
	switch h.metadata.Scene {
	case "query":
		return base.SceneTypeQuery
	default:
		return base.SceneTypeAll
	}
}

func (h *Handler) getEngineType(ctx context.Context) storepb.Engine {
	instanceID := h.getInstanceID()
	if instanceID == "" {
		return storepb.Engine_ENGINE_UNSPECIFIED
	}

	instance, err := h.store.GetInstanceV2(ctx, &store.FindInstanceMessage{
		ResourceID: &instanceID,
	})
	if err != nil {
		slog.Error("Failed to get instance", log.BBError(err))
		return storepb.Engine_ENGINE_UNSPECIFIED
	}
	if instance == nil {
		slog.Error("Instance not found", slog.String("instanceID", instanceID))
		return storepb.Engine_ENGINE_UNSPECIFIED
	}
	return instance.Metadata.GetEngine()
}

func (h *Handler) checkInitialized(req *jsonrpc2.Request) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if Method(req.Method) != LSPMethodInitialize && h.init == nil {
		return errors.New("server must be initialized first")
	}
	return nil
}

func (h *Handler) checkReady() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.shutDown {
		return errors.New("server is shutting down")
	}
	return nil
}

func (h *Handler) reset(params *lsp.InitializeParams) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.init = params
	h.fs = NewMemFS()
	return nil
}

func (h *Handler) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = errors.Errorf("panic: %v", r)
			}
			slog.Error("Panic in LSP handler", log.BBError(err), slog.String("method", req.Method), log.BBStack("panic-stack"))
		}
	}()

	// Handle ping request before checking if the server is initialized.
	if Method(req.Method) == LSPMethodPing {
		return PingResult{Result: "pong"}, nil
	}

	if err := h.checkInitialized(req); err != nil {
		return nil, err
	}
	if err := h.checkReady(); err != nil {
		return nil, err
	}

	switch Method(req.Method) {
	case LSPMethodInitialize:
		if h.init != nil {
			return nil, errors.New("server is already initialized")
		}
		if req.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.InitializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		if err := h.reset(&params); err != nil {
			return nil, err
		}

		return lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				TextDocumentSync: lsp.Incremental,
				CompletionProvider: &lsp.CompletionOptions{
					TriggerCharacters: []string{".", " "},
				},
				ExecuteCommandProvider: &lsp.ExecuteCommandOptions{
					Commands: []string{string(CommandNameSetMetadata)},
				},
			},
		}, nil
	case LSPMethodInitialized:
		// A notification that the client is ready to receive requests. Ignore.
		return nil, nil
	case LSPMethodShutdown:
		h.ShutDown()
		return nil, nil
	case LSPMethodExit:
		conn.Close()
		h.ShutDown()
		return nil, nil
	case LSPMethodCancelRequest:
		var params lsp.CancelParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal cancel request")
		}
		var id jsonrpc2.ID
		if v, ok := params.ID.(string); ok {
			id = jsonrpc2.ID{
				Str:      v,
				IsString: true,
			}
		} else if v, ok := params.ID.(float64); ok {
			// handle json number
			id = jsonrpc2.ID{
				Num: uint64(v),
			}
		}
		cancelFAny, loaded := h.cancelF.LoadAndDelete(id)
		if loaded {
			if cancelF, ok := cancelFAny.(context.CancelFunc); ok {
				cancelF()
			}
		}
		return nil, nil
	case LSPMethodSetTrace:
		// Do nothing for now.
		return nil, nil
	case LSPMethodExecuteCommand:
		if req.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.ExecuteCommandParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		switch CommandName(params.Command) {
		case CommandNameSetMetadata:
			var setMetadataParams SetMetadataCommandParams
			if err := json.Unmarshal(*req.Params, &setMetadataParams); err != nil {
				return nil, err
			}
			if len(setMetadataParams.Arguments) != 1 {
				return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams, Message: "expected exactly one argument"}
			}
			h.setMetadata(setMetadataParams.Arguments[0])
			return nil, nil
		default:
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams, Message: fmt.Sprintf("command not supported: %s", params.Command)}
		}
	case LSPMethodCompletion:
		if req.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.CompletionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		childCtx, cancel := context.WithCancel(ctx)
		h.cancelF.Store(req.ID, cancel)
		defer func() {
			cancel()
			h.cancelF.Delete(req.ID)
		}()
		return h.handleTextDocumentCompletion(childCtx, conn, req, params)
	default:
		if isFileSystemRequest(req.Method) {
			_, _, err := h.handleFileSystemRequest(ctx, conn, req)
			return nil, err
		}
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", req.Method)}
	}
}
