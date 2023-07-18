// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RolloutService_GetPlan_FullMethodName             = "/bytebase.v1.RolloutService/GetPlan"
	RolloutService_ListPlans_FullMethodName           = "/bytebase.v1.RolloutService/ListPlans"
	RolloutService_CreatePlan_FullMethodName          = "/bytebase.v1.RolloutService/CreatePlan"
	RolloutService_UpdatePlan_FullMethodName          = "/bytebase.v1.RolloutService/UpdatePlan"
	RolloutService_GetRollout_FullMethodName          = "/bytebase.v1.RolloutService/GetRollout"
	RolloutService_CreateRollout_FullMethodName       = "/bytebase.v1.RolloutService/CreateRollout"
	RolloutService_PreviewRollout_FullMethodName      = "/bytebase.v1.RolloutService/PreviewRollout"
	RolloutService_ListTaskRuns_FullMethodName        = "/bytebase.v1.RolloutService/ListTaskRuns"
	RolloutService_ListPlanCheckRuns_FullMethodName   = "/bytebase.v1.RolloutService/ListPlanCheckRuns"
	RolloutService_BatchRunTasks_FullMethodName       = "/bytebase.v1.RolloutService/BatchRunTasks"
	RolloutService_BatchSkipTasks_FullMethodName      = "/bytebase.v1.RolloutService/BatchSkipTasks"
	RolloutService_BatchCancelTaskRuns_FullMethodName = "/bytebase.v1.RolloutService/BatchCancelTaskRuns"
)

// RolloutServiceClient is the client API for RolloutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolloutServiceClient interface {
	GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*Plan, error)
	ListPlans(ctx context.Context, in *ListPlansRequest, opts ...grpc.CallOption) (*ListPlansResponse, error)
	CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*Plan, error)
	UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*Plan, error)
	GetRollout(ctx context.Context, in *GetRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	CreateRollout(ctx context.Context, in *CreateRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	PreviewRollout(ctx context.Context, in *PreviewRolloutRequest, opts ...grpc.CallOption) (*Rollout, error)
	ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error)
	ListPlanCheckRuns(ctx context.Context, in *ListPlanCheckRunsRequest, opts ...grpc.CallOption) (*ListPlanCheckRunsResponse, error)
	BatchRunTasks(ctx context.Context, in *BatchRunTasksRequest, opts ...grpc.CallOption) (*BatchRunTasksResponse, error)
	BatchSkipTasks(ctx context.Context, in *BatchSkipTasksRequest, opts ...grpc.CallOption) (*BatchSkipTasksResponse, error)
	BatchCancelTaskRuns(ctx context.Context, in *BatchCancelTaskRunsRequest, opts ...grpc.CallOption) (*BatchCancelTaskRunsResponse, error)
}

type rolloutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolloutServiceClient(cc grpc.ClientConnInterface) RolloutServiceClient {
	return &rolloutServiceClient{cc}
}

func (c *rolloutServiceClient) GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*Plan, error) {
	out := new(Plan)
	err := c.cc.Invoke(ctx, RolloutService_GetPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) ListPlans(ctx context.Context, in *ListPlansRequest, opts ...grpc.CallOption) (*ListPlansResponse, error) {
	out := new(ListPlansResponse)
	err := c.cc.Invoke(ctx, RolloutService_ListPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*Plan, error) {
	out := new(Plan)
	err := c.cc.Invoke(ctx, RolloutService_CreatePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*Plan, error) {
	out := new(Plan)
	err := c.cc.Invoke(ctx, RolloutService_UpdatePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) GetRollout(ctx context.Context, in *GetRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_GetRollout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) CreateRollout(ctx context.Context, in *CreateRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_CreateRollout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) PreviewRollout(ctx context.Context, in *PreviewRolloutRequest, opts ...grpc.CallOption) (*Rollout, error) {
	out := new(Rollout)
	err := c.cc.Invoke(ctx, RolloutService_PreviewRollout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error) {
	out := new(ListTaskRunsResponse)
	err := c.cc.Invoke(ctx, RolloutService_ListTaskRuns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) ListPlanCheckRuns(ctx context.Context, in *ListPlanCheckRunsRequest, opts ...grpc.CallOption) (*ListPlanCheckRunsResponse, error) {
	out := new(ListPlanCheckRunsResponse)
	err := c.cc.Invoke(ctx, RolloutService_ListPlanCheckRuns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchRunTasks(ctx context.Context, in *BatchRunTasksRequest, opts ...grpc.CallOption) (*BatchRunTasksResponse, error) {
	out := new(BatchRunTasksResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchRunTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchSkipTasks(ctx context.Context, in *BatchSkipTasksRequest, opts ...grpc.CallOption) (*BatchSkipTasksResponse, error) {
	out := new(BatchSkipTasksResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchSkipTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolloutServiceClient) BatchCancelTaskRuns(ctx context.Context, in *BatchCancelTaskRunsRequest, opts ...grpc.CallOption) (*BatchCancelTaskRunsResponse, error) {
	out := new(BatchCancelTaskRunsResponse)
	err := c.cc.Invoke(ctx, RolloutService_BatchCancelTaskRuns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolloutServiceServer is the server API for RolloutService service.
// All implementations must embed UnimplementedRolloutServiceServer
// for forward compatibility
type RolloutServiceServer interface {
	GetPlan(context.Context, *GetPlanRequest) (*Plan, error)
	ListPlans(context.Context, *ListPlansRequest) (*ListPlansResponse, error)
	CreatePlan(context.Context, *CreatePlanRequest) (*Plan, error)
	UpdatePlan(context.Context, *UpdatePlanRequest) (*Plan, error)
	GetRollout(context.Context, *GetRolloutRequest) (*Rollout, error)
	CreateRollout(context.Context, *CreateRolloutRequest) (*Rollout, error)
	PreviewRollout(context.Context, *PreviewRolloutRequest) (*Rollout, error)
	ListTaskRuns(context.Context, *ListTaskRunsRequest) (*ListTaskRunsResponse, error)
	ListPlanCheckRuns(context.Context, *ListPlanCheckRunsRequest) (*ListPlanCheckRunsResponse, error)
	BatchRunTasks(context.Context, *BatchRunTasksRequest) (*BatchRunTasksResponse, error)
	BatchSkipTasks(context.Context, *BatchSkipTasksRequest) (*BatchSkipTasksResponse, error)
	BatchCancelTaskRuns(context.Context, *BatchCancelTaskRunsRequest) (*BatchCancelTaskRunsResponse, error)
	mustEmbedUnimplementedRolloutServiceServer()
}

// UnimplementedRolloutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRolloutServiceServer struct {
}

func (UnimplementedRolloutServiceServer) GetPlan(context.Context, *GetPlanRequest) (*Plan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlan not implemented")
}
func (UnimplementedRolloutServiceServer) ListPlans(context.Context, *ListPlansRequest) (*ListPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlans not implemented")
}
func (UnimplementedRolloutServiceServer) CreatePlan(context.Context, *CreatePlanRequest) (*Plan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedRolloutServiceServer) UpdatePlan(context.Context, *UpdatePlanRequest) (*Plan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlan not implemented")
}
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
func (UnimplementedRolloutServiceServer) ListPlanCheckRuns(context.Context, *ListPlanCheckRunsRequest) (*ListPlanCheckRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlanCheckRuns not implemented")
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

// UnsafeRolloutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolloutServiceServer will
// result in compilation errors.
type UnsafeRolloutServiceServer interface {
	mustEmbedUnimplementedRolloutServiceServer()
}

func RegisterRolloutServiceServer(s grpc.ServiceRegistrar, srv RolloutServiceServer) {
	s.RegisterService(&RolloutService_ServiceDesc, srv)
}

func _RolloutService_GetPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).GetPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_GetPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).GetPlan(ctx, req.(*GetPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_ListPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).ListPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_ListPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).ListPlans(ctx, req.(*ListPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_CreatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).CreatePlan(ctx, req.(*CreatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolloutService_UpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).UpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_UpdatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).UpdatePlan(ctx, req.(*UpdatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _RolloutService_ListPlanCheckRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlanCheckRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolloutServiceServer).ListPlanCheckRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolloutService_ListPlanCheckRuns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolloutServiceServer).ListPlanCheckRuns(ctx, req.(*ListPlanCheckRunsRequest))
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
			MethodName: "GetPlan",
			Handler:    _RolloutService_GetPlan_Handler,
		},
		{
			MethodName: "ListPlans",
			Handler:    _RolloutService_ListPlans_Handler,
		},
		{
			MethodName: "CreatePlan",
			Handler:    _RolloutService_CreatePlan_Handler,
		},
		{
			MethodName: "UpdatePlan",
			Handler:    _RolloutService_UpdatePlan_Handler,
		},
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
			MethodName: "ListPlanCheckRuns",
			Handler:    _RolloutService_ListPlanCheckRuns_Handler,
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
