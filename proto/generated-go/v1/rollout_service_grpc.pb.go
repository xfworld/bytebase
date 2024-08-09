// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: v1/rollout_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RolloutService_GetRollout_FullMethodName          = "/bytebase.v1.RolloutService/GetRollout"
	RolloutService_CreateRollout_FullMethodName       = "/bytebase.v1.RolloutService/CreateRollout"
	RolloutService_PreviewRollout_FullMethodName      = "/bytebase.v1.RolloutService/PreviewRollout"
	RolloutService_ListTaskRuns_FullMethodName        = "/bytebase.v1.RolloutService/ListTaskRuns"
	RolloutService_GetTaskRunLog_FullMethodName       = "/bytebase.v1.RolloutService/GetTaskRunLog"
	RolloutService_GetTaskRunSession_FullMethodName   = "/bytebase.v1.RolloutService/GetTaskRunSession"
	RolloutService_BatchRunTasks_FullMethodName       = "/bytebase.v1.RolloutService/BatchRunTasks"
	RolloutService_BatchSkipTasks_FullMethodName      = "/bytebase.v1.RolloutService/BatchSkipTasks"
	RolloutService_BatchCancelTaskRuns_FullMethodName = "/bytebase.v1.RolloutService/BatchCancelTaskRuns"
)

// RolloutServiceClient is the client API for RolloutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolloutServiceClient interface {
	GetRollout(ctx context.Context, in *GetRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	CreateRollout(ctx context.Context, in *CreateRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	PreviewRollout(ctx context.Context, in *PreviewRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error)
	GetTaskRunLog(ctx context.Context, in *GetTaskRunLogRequest, opts ...grpc.CallOption) (*TaskRunLog, error)
	GetTaskRunSession(ctx context.Context, in *GetTaskRunSessionRequest, opts ...grpc.CallOption) (*TaskRunSession, error)
	// BatchRunTasks creates task runs for the specified tasks.
	// DataExport issue only allows the creator to run the task.
	// Users with "bb.taskRuns.create" permission can run the task, e.g. Workspace Admin and DBA.
	// Follow role-based rollout policy for the environment.
	BatchRunTasks(ctx context.Context, in *BatchRunTasksRequest, opts ...grpc.CallOption) (*BatchRunTasksResponse, error)
	// BatchSkipTasks skips the specified tasks.
	// The access is the same as BatchRunTasks().
	BatchSkipTasks(ctx context.Context, in *BatchSkipTasksRequest, opts ...grpc.CallOption) (*BatchSkipTasksResponse, error)
	// BatchSkipTasks cancels the specified task runs in batch.
	// The access is the same as BatchRunTasks().
	BatchCancelTaskRuns(ctx context.Context, in *BatchCancelTaskRunsRequest, opts ...grpc.CallOption) (*BatchCancelTaskRunsResponse, error)
}

type rolloutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolloutServiceClient(cc grpc.ClientConnInterface) RolloutServiceClient {
	return &rolloutServiceClient{cc}
}

func (c *rolloutServiceClient) GetRollout(ctx context.Context, in *GetRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_GetRollout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) CreateRollout(ctx context.Context, in *CreateRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_CreateRollout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) PreviewRollout(ctx context.Context, in *PreviewRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_PreviewRollout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTaskRunsResponse)
	err := c.cc.Invoke(ctx, RolloutService_ListTaskRuns_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) GetTaskRunLog(ctx context.Context, in *GetTaskRunLogRequest, opts ...grpc.CallOption) (*TaskRunLog, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskRunLog)
	err := c.cc.Invoke(ctx, RolloutService_GetTaskRunLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) GetTaskRunSession(ctx context.Context, in *GetTaskRunSessionRequest, opts ...grpc.CallOption) (*TaskRunSession, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskRunSession)
	err := c.cc.Invoke(ctx, RolloutService_GetTaskRunSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchRunTasks(ctx context.Context, in *BatchRunTasksRequest, opts ...grpc.CallOption) (*BatchRunTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchRunTasksResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchRunTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchSkipTasks(ctx context.Context, in *BatchSkipTasksRequest, opts ...grpc.CallOption) (*BatchSkipTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchSkipTasksResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchSkipTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchCancelTaskRuns(ctx context.Context, in *BatchCancelTaskRunsRequest, opts ...grpc.CallOption) (*BatchCancelTaskRunsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchCancelTaskRunsResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchCancelTaskRuns_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolloutServiceServer is the server API for RolloutService service.
// All implementations must embed UnimplementedRolloutServiceServer
// for forward compatibility.
type RolloutServiceServer interface {
	GetRollout(context.Context, *GetRolloutRequest) (*Rollout, error)
	CreateRollout(context.Context, *CreateRolloutRequest) (*Rollout, error)
	PreviewRollout(context.Context, *PreviewRolloutRequest) (*Rollout, error)
	ListTaskRuns(context.Context, *ListTaskRunsRequest) (*ListTaskRunsResponse, error)
	GetTaskRunLog(context.Context, *GetTaskRunLogRequest) (*TaskRunLog, error)
	GetTaskRunSession(context.Context, *GetTaskRunSessionRequest) (*TaskRunSession, error)
	// BatchRunTasks creates task runs for the specified tasks.
	// DataExport issue only allows the creator to run the task.
	// Users with "bb.taskRuns.create" permission can run the task, e.g. Workspace Admin and DBA.
	// Follow role-based rollout policy for the environment.
	BatchRunTasks(context.Context, *BatchRunTasksRequest) (*BatchRunTasksResponse, error)
	// BatchSkipTasks skips the specified tasks.
	// The access is the same as BatchRunTasks().
	BatchSkipTasks(context.Context, *BatchSkipTasksRequest) (*BatchSkipTasksResponse, error)
	// BatchSkipTasks cancels the specified task runs in batch.
	// The access is the same as BatchRunTasks().
	BatchCancelTaskRuns(context.Context, *BatchCancelTaskRunsRequest) (*BatchCancelTaskRunsResponse, error)
	mustEmbedUnimplementedRolloutServiceServer()
}

// UnimplementedRolloutServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRolloutServiceServer struct{}

func (UnimplementedRolloutServiceServer) GetRollout(context.Context, *GetRolloutRequest) (*Rollout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRollout not implemented")
}
func (UnimplementedRolloutServiceServer) CreateRollout(context.Context, *CreateRolloutRequest) (*Rollout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRollout not implemented")
}
func (UnimplementedRolloutServiceServer) PreviewRollout(context.Context, *PreviewRolloutRequest) (*Rollout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRollout not implemented")
}
func (UnimplementedRolloutServiceServer) ListTaskRuns(context.Context, *ListTaskRunsRequest) (*ListTaskRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskRuns not implemented")
}
func (UnimplementedRolloutServiceServer) GetTaskRunLog(context.Context, *GetTaskRunLogRequest) (*TaskRunLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskRunLog not implemented")
}
func (UnimplementedRolloutServiceServer) GetTaskRunSession(context.Context, *GetTaskRunSessionRequest) (*TaskRunSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskRunSession not implemented")
}
func (UnimplementedRolloutServiceServer) BatchRunTasks(context.Context, *BatchRunTasksRequest) (*BatchRunTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchRunTasks not implemented")
}
func (UnimplementedRolloutServiceServer) BatchSkipTasks(context.Context, *BatchSkipTasksRequest) (*BatchSkipTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSkipTasks not implemented")
}
func (UnimplementedRolloutServiceServer) BatchCancelTaskRuns(context.Context, *BatchCancelTaskRunsRequest) (*BatchCancelTaskRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCancelTaskRuns not implemented")
}
func (UnimplementedRolloutServiceServer) mustEmbedUnimplementedRolloutServiceServer() {}
func (UnimplementedRolloutServiceServer) testEmbeddedByValue()                        {}

// UnsafeRolloutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolloutServiceServer will
// result in compilation errors.
type UnsafeRolloutServiceServer interface {
	mustEmbedUnimplementedRolloutServiceServer()
}

func RegisterRolloutServiceServer(s grpc.ServiceRegistrar, srv RolloutServiceServer) {
	// If the following call pancis, it indicates UnimplementedRolloutServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RolloutService_ServiceDesc, srv)
}

func _RolloutService_GetRollout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolloutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).GetRollout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_GetRollout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).GetRollout(ctx, req.(*GetRolloutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_CreateRollout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRolloutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).CreateRollout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_CreateRollout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).CreateRollout(ctx, req.(*CreateRolloutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_PreviewRollout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewRolloutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).PreviewRollout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_PreviewRollout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).PreviewRollout(ctx, req.(*PreviewRolloutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_ListTaskRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).ListTaskRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_ListTaskRuns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).ListTaskRuns(ctx, req.(*ListTaskRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_GetTaskRunLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRunLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).GetTaskRunLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_GetTaskRunLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).GetTaskRunLog(ctx, req.(*GetTaskRunLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_GetTaskRunSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRunSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).GetTaskRunSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_GetTaskRunSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).GetTaskRunSession(ctx, req.(*GetTaskRunSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_BatchRunTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRunTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).BatchRunTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_BatchRunTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).BatchRunTasks(ctx, req.(*BatchRunTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_BatchSkipTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSkipTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).BatchSkipTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_BatchSkipTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).BatchSkipTasks(ctx, req.(*BatchSkipTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_BatchCancelTaskRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCancelTaskRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).BatchCancelTaskRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_BatchCancelTaskRuns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).BatchCancelTaskRuns(ctx, req.(*BatchCancelTaskRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RolloutService_ServiceDesc is the grpc.ServiceDesc for RolloutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RolloutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.RolloutService",
	HandlerType: (*RolloutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRollout",
			Handler:    _RolloutService_GetRollout_Handler,
		},
		{
			MethodName: "CreateRollout",
			Handler:    _RolloutService_CreateRollout_Handler,
		},
		{
			MethodName: "PreviewRollout",
			Handler:    _RolloutService_PreviewRollout_Handler,
		},
		{
			MethodName: "ListTaskRuns",
			Handler:    _RolloutService_ListTaskRuns_Handler,
		},
		{
			MethodName: "GetTaskRunLog",
			Handler:    _RolloutService_GetTaskRunLog_Handler,
		},
		{
			MethodName: "GetTaskRunSession",
			Handler:    _RolloutService_GetTaskRunSession_Handler,
		},
		{
			MethodName: "BatchRunTasks",
			Handler:    _RolloutService_BatchRunTasks_Handler,
		},
		{
			MethodName: "BatchSkipTasks",
			Handler:    _RolloutService_BatchSkipTasks_Handler,
		},
		{
			MethodName: "BatchCancelTaskRuns",
			Handler:    _RolloutService_BatchCancelTaskRuns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/rollout_service.proto",
}
