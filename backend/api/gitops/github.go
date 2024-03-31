package gitops

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/common/log"
	"github.com/bytebase/bytebase/backend/plugin/vcs"
	"github.com/bytebase/bytebase/backend/plugin/vcs/github"
	"github.com/bytebase/bytebase/backend/store"
)

const (
	closeAction = "closed"
)

func getGitHubPullRequestInfo(ctx context.Context, vcsProvider *store.VCSProviderMessage, vcsConnector *store.VCSConnectorMessage, body []byte) (*pullRequestInfo, error) {
	var pushEvent github.PullRequestPushEvent
	if err := json.Unmarshal(body, &pushEvent); err != nil {
		return nil, errors.Errorf("failed to unmarshal push event, error %v", err)
	}
	if pushEvent.Action != closeAction {
		return nil, errors.Errorf("invalid webhook event action, got %s, want closed", pushEvent.Action)
	}

	if pushEvent.PullRequest.Base.Ref != vcsConnector.Payload.Branch {
		return nil, errors.Errorf("committed to branch %q, want branch %q", pushEvent.PullRequest.Base.Ref, vcsConnector.Payload.Branch)
	}
	oauthContext := &common.OauthContext{
		AccessToken: vcsProvider.AccessToken,
	}

	mrFiles, err := vcs.Get(vcs.GitHub, vcs.ProviderConfig{}).ListPullRequestFile(ctx, oauthContext, vcsProvider.InstanceURL, vcsConnector.Payload.ExternalId, fmt.Sprintf("%d", pushEvent.Number))
	if err != nil {
		return nil, errors.Errorf("failed to list merge %q request files, error %v", pushEvent.PullRequest.URL, err)
	}

	prInfo := &pullRequestInfo{
		// email. How do we determine the user for GitHub user?
		url:         pushEvent.PullRequest.URL,
		title:       pushEvent.PullRequest.Title,
		description: pushEvent.PullRequest.Body,
	}
	for _, v := range mrFiles {
		if v.IsDeleted {
			continue
		}
		if filepath.Dir(v.Path) != vcsConnector.Payload.BaseDirectory {
			continue
		}
		change, err := getFileChange(v.Path)
		if err != nil {
			slog.Error("failed to get file change info", slog.String("path", v.Path), log.BBError(err))
		}
		if change != nil {
			change.path = v.Path
			prInfo.changes = append(prInfo.changes, change)
		}
	}
	for _, file := range prInfo.changes {
		content, err := vcs.Get(vcs.GitHub, vcs.ProviderConfig{}).ReadFileContent(ctx, oauthContext, vcsProvider.InstanceURL, vcsConnector.Payload.ExternalId, file.path, vcs.RefInfo{RefType: vcs.RefTypeCommit, RefName: pushEvent.PullRequest.Head.SHA})
		if err != nil {
			return nil, errors.Errorf("failed read file content, merge request %q, file %q, error %v", pushEvent.PullRequest.URL, file.path, err)
		}
		file.content = content
	}
	return prInfo, nil
}