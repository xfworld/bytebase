// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: v1/branch_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BranchService_GetBranch_FullMethodName    = "/bytebase.v1.BranchService/GetBranch"
	BranchService_ListBranches_FullMethodName = "/bytebase.v1.BranchService/ListBranches"
	BranchService_CreateBranch_FullMethodName = "/bytebase.v1.BranchService/CreateBranch"
	BranchService_UpdateBranch_FullMethodName = "/bytebase.v1.BranchService/UpdateBranch"
	BranchService_MergeBranch_FullMethodName  = "/bytebase.v1.BranchService/MergeBranch"
	BranchService_RebaseBranch_FullMethodName = "/bytebase.v1.BranchService/RebaseBranch"
	BranchService_DeleteBranch_FullMethodName = "/bytebase.v1.BranchService/DeleteBranch"
	BranchService_DiffDatabase_FullMethodName = "/bytebase.v1.BranchService/DiffDatabase"
	BranchService_DiffMetadata_FullMethodName = "/bytebase.v1.BranchService/DiffMetadata"
)

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchServiceClient interface {
	GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (*ListBranchesResponse, error)
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	MergeBranch(ctx context.Context, in *MergeBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	RebaseBranch(ctx context.Context, in *RebaseBranchRequest, opts ...grpc.CallOption) (*RebaseBranchResponse, error)
	DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DiffDatabase works similar to branch rebase.
	// 1) set the base as the schema of a database;
	// 2) apply the changes between base and head of branch to the new base (schema of database);
	// 3) return the diff DDLs similar to DiffSchema in database service.
	// 4) return the conflict schema if conflict needs to be resolved by user. Once resolved, user
	// will call DiffSchema() in database service to get diff DDLs.
	DiffDatabase(ctx context.Context, in *DiffDatabaseRequest, opts ...grpc.CallOption) (*DiffDatabaseResponse, error)
	DiffMetadata(ctx context.Context, in *DiffMetadataRequest, opts ...grpc.CallOption) (*DiffMetadataResponse, error)
}

type branchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchServiceClient(cc grpc.ClientConnInterface) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Branch)
	err := c.cc.Invoke(ctx, BranchService_GetBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) ListBranches(ctx context.Context, in *ListBranchesRequest, opts ...grpc.CallOption) (*ListBranchesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBranchesResponse)
	err := c.cc.Invoke(ctx, BranchService_ListBranches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Branch)
	err := c.cc.Invoke(ctx, BranchService_CreateBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Branch)
	err := c.cc.Invoke(ctx, BranchService_UpdateBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) MergeBranch(ctx context.Context, in *MergeBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Branch)
	err := c.cc.Invoke(ctx, BranchService_MergeBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) RebaseBranch(ctx context.Context, in *RebaseBranchRequest, opts ...grpc.CallOption) (*RebaseBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RebaseBranchResponse)
	err := c.cc.Invoke(ctx, BranchService_RebaseBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BranchService_DeleteBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) DiffDatabase(ctx context.Context, in *DiffDatabaseRequest, opts ...grpc.CallOption) (*DiffDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiffDatabaseResponse)
	err := c.cc.Invoke(ctx, BranchService_DiffDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) DiffMetadata(ctx context.Context, in *DiffMetadataRequest, opts ...grpc.CallOption) (*DiffMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiffMetadataResponse)
	err := c.cc.Invoke(ctx, BranchService_DiffMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchServiceServer is the server API for BranchService service.
// All implementations must embed UnimplementedBranchServiceServer
// for forward compatibility.
type BranchServiceServer interface {
	GetBranch(context.Context, *GetBranchRequest) (*Branch, error)
	ListBranches(context.Context, *ListBranchesRequest) (*ListBranchesResponse, error)
	CreateBranch(context.Context, *CreateBranchRequest) (*Branch, error)
	UpdateBranch(context.Context, *UpdateBranchRequest) (*Branch, error)
	MergeBranch(context.Context, *MergeBranchRequest) (*Branch, error)
	RebaseBranch(context.Context, *RebaseBranchRequest) (*RebaseBranchResponse, error)
	DeleteBranch(context.Context, *DeleteBranchRequest) (*emptypb.Empty, error)
	// DiffDatabase works similar to branch rebase.
	// 1) set the base as the schema of a database;
	// 2) apply the changes between base and head of branch to the new base (schema of database);
	// 3) return the diff DDLs similar to DiffSchema in database service.
	// 4) return the conflict schema if conflict needs to be resolved by user. Once resolved, user
	// will call DiffSchema() in database service to get diff DDLs.
	DiffDatabase(context.Context, *DiffDatabaseRequest) (*DiffDatabaseResponse, error)
	DiffMetadata(context.Context, *DiffMetadataRequest) (*DiffMetadataResponse, error)
	mustEmbedUnimplementedBranchServiceServer()
}

// UnimplementedBranchServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBranchServiceServer struct{}

func (UnimplementedBranchServiceServer) GetBranch(context.Context, *GetBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranch not implemented")
}
func (UnimplementedBranchServiceServer) ListBranches(context.Context, *ListBranchesRequest) (*ListBranchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBranches not implemented")
}
func (UnimplementedBranchServiceServer) CreateBranch(context.Context, *CreateBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (UnimplementedBranchServiceServer) UpdateBranch(context.Context, *UpdateBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBranch not implemented")
}
func (UnimplementedBranchServiceServer) MergeBranch(context.Context, *MergeBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeBranch not implemented")
}
func (UnimplementedBranchServiceServer) RebaseBranch(context.Context, *RebaseBranchRequest) (*RebaseBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebaseBranch not implemented")
}
func (UnimplementedBranchServiceServer) DeleteBranch(context.Context, *DeleteBranchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranch not implemented")
}
func (UnimplementedBranchServiceServer) DiffDatabase(context.Context, *DiffDatabaseRequest) (*DiffDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffDatabase not implemented")
}
func (UnimplementedBranchServiceServer) DiffMetadata(context.Context, *DiffMetadataRequest) (*DiffMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffMetadata not implemented")
}
func (UnimplementedBranchServiceServer) mustEmbedUnimplementedBranchServiceServer() {}
func (UnimplementedBranchServiceServer) testEmbeddedByValue()                       {}

// UnsafeBranchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchServiceServer will
// result in compilation errors.
type UnsafeBranchServiceServer interface {
	mustEmbedUnimplementedBranchServiceServer()
}

func RegisterBranchServiceServer(s grpc.ServiceRegistrar, srv BranchServiceServer) {
	// If the following call pancis, it indicates UnimplementedBranchServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BranchService_ServiceDesc, srv)
}

func _BranchService_GetBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_GetBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranch(ctx, req.(*GetBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_ListBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBranchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).ListBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_ListBranches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).ListBranches(ctx, req.(*ListBranchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_CreateBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_UpdateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).UpdateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_UpdateBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).UpdateBranch(ctx, req.(*UpdateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_MergeBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).MergeBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_MergeBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).MergeBranch(ctx, req.(*MergeBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_RebaseBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebaseBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).RebaseBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_RebaseBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).RebaseBranch(ctx, req.(*RebaseBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_DeleteBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).DeleteBranch(ctx, req.(*DeleteBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_DiffDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).DiffDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_DiffDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).DiffDatabase(ctx, req.(*DiffDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_DiffMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).DiffMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_DiffMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).DiffMetadata(ctx, req.(*DiffMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchService_ServiceDesc is the grpc.ServiceDesc for BranchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBranch",
			Handler:    _BranchService_GetBranch_Handler,
		},
		{
			MethodName: "ListBranches",
			Handler:    _BranchService_ListBranches_Handler,
		},
		{
			MethodName: "CreateBranch",
			Handler:    _BranchService_CreateBranch_Handler,
		},
		{
			MethodName: "UpdateBranch",
			Handler:    _BranchService_UpdateBranch_Handler,
		},
		{
			MethodName: "MergeBranch",
			Handler:    _BranchService_MergeBranch_Handler,
		},
		{
			MethodName: "RebaseBranch",
			Handler:    _BranchService_RebaseBranch_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _BranchService_DeleteBranch_Handler,
		},
		{
			MethodName: "DiffDatabase",
			Handler:    _BranchService_DiffDatabase_Handler,
		},
		{
			MethodName: "DiffMetadata",
			Handler:    _BranchService_DiffMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/branch_service.proto",
}
