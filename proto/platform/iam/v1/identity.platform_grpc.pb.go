// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: identity.platform.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Identities_Create_FullMethodName = "/chainguard.platform.iam.Identities/Create"
	Identities_Update_FullMethodName = "/chainguard.platform.iam.Identities/Update"
	Identities_List_FullMethodName   = "/chainguard.platform.iam.Identities/List"
	Identities_Lookup_FullMethodName = "/chainguard.platform.iam.Identities/Lookup"
	Identities_Delete_FullMethodName = "/chainguard.platform.iam.Identities/Delete"
)

// IdentitiesClient is the client API for Identities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentitiesClient interface {
	Create(ctx context.Context, in *CreateIdentityRequest, opts ...grpc.CallOption) (*Identity, error)
	Update(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error)
	List(ctx context.Context, in *IdentityFilter, opts ...grpc.CallOption) (*IdentityList, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Identity, error)
	Delete(ctx context.Context, in *DeleteIdentityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentitiesClient(cc grpc.ClientConnInterface) IdentitiesClient {
	return &identitiesClient{cc}
}

func (c *identitiesClient) Create(ctx context.Context, in *CreateIdentityRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, Identities_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitiesClient) Update(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, Identities_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitiesClient) List(ctx context.Context, in *IdentityFilter, opts ...grpc.CallOption) (*IdentityList, error) {
	out := new(IdentityList)
	err := c.cc.Invoke(ctx, Identities_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitiesClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, Identities_Lookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitiesClient) Delete(ctx context.Context, in *DeleteIdentityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Identities_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentitiesServer is the server API for Identities service.
// All implementations must embed UnimplementedIdentitiesServer
// for forward compatibility
type IdentitiesServer interface {
	Create(context.Context, *CreateIdentityRequest) (*Identity, error)
	Update(context.Context, *Identity) (*Identity, error)
	List(context.Context, *IdentityFilter) (*IdentityList, error)
	Lookup(context.Context, *LookupRequest) (*Identity, error)
	Delete(context.Context, *DeleteIdentityRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentitiesServer()
}

// UnimplementedIdentitiesServer must be embedded to have forward compatible implementations.
type UnimplementedIdentitiesServer struct {
}

func (UnimplementedIdentitiesServer) Create(context.Context, *CreateIdentityRequest) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIdentitiesServer) Update(context.Context, *Identity) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIdentitiesServer) List(context.Context, *IdentityFilter) (*IdentityList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIdentitiesServer) Lookup(context.Context, *LookupRequest) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedIdentitiesServer) Delete(context.Context, *DeleteIdentityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIdentitiesServer) mustEmbedUnimplementedIdentitiesServer() {}

// UnsafeIdentitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentitiesServer will
// result in compilation errors.
type UnsafeIdentitiesServer interface {
	mustEmbedUnimplementedIdentitiesServer()
}

func RegisterIdentitiesServer(s grpc.ServiceRegistrar, srv IdentitiesServer) {
	s.RegisterService(&Identities_ServiceDesc, srv)
}

func _Identities_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitiesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identities_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitiesServer).Create(ctx, req.(*CreateIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identities_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitiesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identities_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitiesServer).Update(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identities_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitiesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identities_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitiesServer).List(ctx, req.(*IdentityFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identities_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitiesServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identities_Lookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitiesServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identities_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitiesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identities_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitiesServer).Delete(ctx, req.(*DeleteIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Identities_ServiceDesc is the grpc.ServiceDesc for Identities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.iam.Identities",
	HandlerType: (*IdentitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Identities_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Identities_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Identities_List_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Identities_Lookup_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Identities_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.platform.proto",
}
