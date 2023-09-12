// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: oidc.platform.proto

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
	SecurityTokenService_Exchange_FullMethodName = "/chainguard.platform.oidc.SecurityTokenService/Exchange"
)

// SecurityTokenServiceClient is the client API for SecurityTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityTokenServiceClient interface {
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*RawToken, error)
}

type securityTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityTokenServiceClient(cc grpc.ClientConnInterface) SecurityTokenServiceClient {
	return &securityTokenServiceClient{cc}
}

func (c *securityTokenServiceClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*RawToken, error) {
	out := new(RawToken)
	err := c.cc.Invoke(ctx, SecurityTokenService_Exchange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityTokenServiceServer is the server API for SecurityTokenService service.
// All implementations must embed UnimplementedSecurityTokenServiceServer
// for forward compatibility
type SecurityTokenServiceServer interface {
	Exchange(context.Context, *ExchangeRequest) (*RawToken, error)
	mustEmbedUnimplementedSecurityTokenServiceServer()
}

// UnimplementedSecurityTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityTokenServiceServer struct {
}

func (UnimplementedSecurityTokenServiceServer) Exchange(context.Context, *ExchangeRequest) (*RawToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedSecurityTokenServiceServer) mustEmbedUnimplementedSecurityTokenServiceServer() {}

// UnsafeSecurityTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityTokenServiceServer will
// result in compilation errors.
type UnsafeSecurityTokenServiceServer interface {
	mustEmbedUnimplementedSecurityTokenServiceServer()
}

func RegisterSecurityTokenServiceServer(s grpc.ServiceRegistrar, srv SecurityTokenServiceServer) {
	s.RegisterService(&SecurityTokenService_ServiceDesc, srv)
}

func _SecurityTokenService_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityTokenServiceServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityTokenService_Exchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityTokenServiceServer).Exchange(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityTokenService_ServiceDesc is the grpc.ServiceDesc for SecurityTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.oidc.SecurityTokenService",
	HandlerType: (*SecurityTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exchange",
			Handler:    _SecurityTokenService_Exchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oidc.platform.proto",
}
