// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: sts.proto

package pb

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
	StsRpc_GenCosSts_FullMethodName      = "/sts.sts_rpc/genCosSts"
	StsRpc_GenSignedUrl_FullMethodName   = "/sts.sts_rpc/genSignedUrl"
	StsRpc_DeleteObject_FullMethodName   = "/sts.sts_rpc/deleteObject"
	StsRpc_GetAccessToken_FullMethodName = "/sts.sts_rpc/GetAccessToken"
)

// StsRpcClient is the client API for StsRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StsRpcClient interface {
	GenCosSts(ctx context.Context, in *GenCosStsReq, opts ...grpc.CallOption) (*GenCosStsResp, error)
	GenSignedUrl(ctx context.Context, in *GenSignedUrlReq, opts ...grpc.CallOption) (*GenSignedUrlResp, error)
	DeleteObject(ctx context.Context, in *DeleteObjectReq, opts ...grpc.CallOption) (*DeleteObjectResp, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...grpc.CallOption) (*GetAccessTokenResp, error)
}

type stsRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewStsRpcClient(cc grpc.ClientConnInterface) StsRpcClient {
	return &stsRpcClient{cc}
}

func (c *stsRpcClient) GenCosSts(ctx context.Context, in *GenCosStsReq, opts ...grpc.CallOption) (*GenCosStsResp, error) {
	out := new(GenCosStsResp)
	err := c.cc.Invoke(ctx, StsRpc_GenCosSts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stsRpcClient) GenSignedUrl(ctx context.Context, in *GenSignedUrlReq, opts ...grpc.CallOption) (*GenSignedUrlResp, error) {
	out := new(GenSignedUrlResp)
	err := c.cc.Invoke(ctx, StsRpc_GenSignedUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stsRpcClient) DeleteObject(ctx context.Context, in *DeleteObjectReq, opts ...grpc.CallOption) (*DeleteObjectResp, error) {
	out := new(DeleteObjectResp)
	err := c.cc.Invoke(ctx, StsRpc_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stsRpcClient) GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...grpc.CallOption) (*GetAccessTokenResp, error) {
	out := new(GetAccessTokenResp)
	err := c.cc.Invoke(ctx, StsRpc_GetAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StsRpcServer is the server API for StsRpc service.
// All implementations must embed UnimplementedStsRpcServer
// for forward compatibility
type StsRpcServer interface {
	GenCosSts(context.Context, *GenCosStsReq) (*GenCosStsResp, error)
	GenSignedUrl(context.Context, *GenSignedUrlReq) (*GenSignedUrlResp, error)
	DeleteObject(context.Context, *DeleteObjectReq) (*DeleteObjectResp, error)
	GetAccessToken(context.Context, *GetAccessTokenReq) (*GetAccessTokenResp, error)
	mustEmbedUnimplementedStsRpcServer()
}

// UnimplementedStsRpcServer must be embedded to have forward compatible implementations.
type UnimplementedStsRpcServer struct {
}

func (UnimplementedStsRpcServer) GenCosSts(context.Context, *GenCosStsReq) (*GenCosStsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenCosSts not implemented")
}
func (UnimplementedStsRpcServer) GenSignedUrl(context.Context, *GenSignedUrlReq) (*GenSignedUrlResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenSignedUrl not implemented")
}
func (UnimplementedStsRpcServer) DeleteObject(context.Context, *DeleteObjectReq) (*DeleteObjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedStsRpcServer) GetAccessToken(context.Context, *GetAccessTokenReq) (*GetAccessTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedStsRpcServer) mustEmbedUnimplementedStsRpcServer() {}

// UnsafeStsRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StsRpcServer will
// result in compilation errors.
type UnsafeStsRpcServer interface {
	mustEmbedUnimplementedStsRpcServer()
}

func RegisterStsRpcServer(s grpc.ServiceRegistrar, srv StsRpcServer) {
	s.RegisterService(&StsRpc_ServiceDesc, srv)
}

func _StsRpc_GenCosSts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenCosStsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsRpcServer).GenCosSts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StsRpc_GenCosSts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsRpcServer).GenCosSts(ctx, req.(*GenCosStsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StsRpc_GenSignedUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenSignedUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsRpcServer).GenSignedUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StsRpc_GenSignedUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsRpcServer).GenSignedUrl(ctx, req.(*GenSignedUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StsRpc_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsRpcServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StsRpc_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsRpcServer).DeleteObject(ctx, req.(*DeleteObjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StsRpc_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsRpcServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StsRpc_GetAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsRpcServer).GetAccessToken(ctx, req.(*GetAccessTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StsRpc_ServiceDesc is the grpc.ServiceDesc for StsRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StsRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sts.sts_rpc",
	HandlerType: (*StsRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "genCosSts",
			Handler:    _StsRpc_GenCosSts_Handler,
		},
		{
			MethodName: "genSignedUrl",
			Handler:    _StsRpc_GenSignedUrl_Handler,
		},
		{
			MethodName: "deleteObject",
			Handler:    _StsRpc_DeleteObject_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _StsRpc_GetAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sts.proto",
}
