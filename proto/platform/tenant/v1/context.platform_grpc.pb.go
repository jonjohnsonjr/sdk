// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: context.platform.proto

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
	RecordContexts_List_FullMethodName = "/chainguard.platform.tenant.RecordContexts/List"
)

// RecordContextsClient is the client API for RecordContexts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordContextsClient interface {
	List(ctx context.Context, in *RecordContextFilter, opts ...grpc.CallOption) (*RecordContextList, error)
}

type recordContextsClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordContextsClient(cc grpc.ClientConnInterface) RecordContextsClient {
	return &recordContextsClient{cc}
}

func (c *recordContextsClient) List(ctx context.Context, in *RecordContextFilter, opts ...grpc.CallOption) (*RecordContextList, error) {
	out := new(RecordContextList)
	err := c.cc.Invoke(ctx, RecordContexts_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordContextsServer is the server API for RecordContexts service.
// All implementations must embed UnimplementedRecordContextsServer
// for forward compatibility
type RecordContextsServer interface {
	List(context.Context, *RecordContextFilter) (*RecordContextList, error)
	mustEmbedUnimplementedRecordContextsServer()
}

// UnimplementedRecordContextsServer must be embedded to have forward compatible implementations.
type UnimplementedRecordContextsServer struct {
}

func (UnimplementedRecordContextsServer) List(context.Context, *RecordContextFilter) (*RecordContextList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRecordContextsServer) mustEmbedUnimplementedRecordContextsServer() {}

// UnsafeRecordContextsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordContextsServer will
// result in compilation errors.
type UnsafeRecordContextsServer interface {
	mustEmbedUnimplementedRecordContextsServer()
}

func RegisterRecordContextsServer(s grpc.ServiceRegistrar, srv RecordContextsServer) {
	s.RegisterService(&RecordContexts_ServiceDesc, srv)
}

func _RecordContexts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordContextFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordContextsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordContexts_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordContextsServer).List(ctx, req.(*RecordContextFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordContexts_ServiceDesc is the grpc.ServiceDesc for RecordContexts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordContexts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.tenant.RecordContexts",
	HandlerType: (*RecordContextsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RecordContexts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "context.platform.proto",
}
