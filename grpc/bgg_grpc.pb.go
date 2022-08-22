// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: grpc/bgg.proto

package grpc

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

// BGGServiceClient is the client API for BGGService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BGGServiceClient interface {
	FindBoardgameByName(ctx context.Context, in *FindBoardgameByNameRequest, opts ...grpc.CallOption) (*FindBoardgameByNameResponse, error)
	FindBoardgameById(ctx context.Context, in *FindBoardgameByIdRequest, opts ...grpc.CallOption) (*FindBoardgameByIdResponse, error)
	InsertNewBoardGame(ctx context.Context, in *InsertNewBoardgameRequest, opts ...grpc.CallOption) (*InsertNewBoardgameResponse, error)
}

type bGGServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBGGServiceClient(cc grpc.ClientConnInterface) BGGServiceClient {
	return &bGGServiceClient{cc}
}

func (c *bGGServiceClient) FindBoardgameByName(ctx context.Context, in *FindBoardgameByNameRequest, opts ...grpc.CallOption) (*FindBoardgameByNameResponse, error) {
	out := new(FindBoardgameByNameResponse)
	err := c.cc.Invoke(ctx, "/grpc.BGGService/FindBoardgameByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGGServiceClient) FindBoardgameById(ctx context.Context, in *FindBoardgameByIdRequest, opts ...grpc.CallOption) (*FindBoardgameByIdResponse, error) {
	out := new(FindBoardgameByIdResponse)
	err := c.cc.Invoke(ctx, "/grpc.BGGService/FindBoardgameById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGGServiceClient) InsertNewBoardGame(ctx context.Context, in *InsertNewBoardgameRequest, opts ...grpc.CallOption) (*InsertNewBoardgameResponse, error) {
	out := new(InsertNewBoardgameResponse)
	err := c.cc.Invoke(ctx, "/grpc.BGGService/InsertNewBoardGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BGGServiceServer is the server API for BGGService service.
// All implementations must embed UnimplementedBGGServiceServer
// for forward compatibility
type BGGServiceServer interface {
	FindBoardgameByName(context.Context, *FindBoardgameByNameRequest) (*FindBoardgameByNameResponse, error)
	FindBoardgameById(context.Context, *FindBoardgameByIdRequest) (*FindBoardgameByIdResponse, error)
	InsertNewBoardGame(context.Context, *InsertNewBoardgameRequest) (*InsertNewBoardgameResponse, error)
	mustEmbedUnimplementedBGGServiceServer()
}

// UnimplementedBGGServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBGGServiceServer struct {
}

func (UnimplementedBGGServiceServer) FindBoardgameByName(context.Context, *FindBoardgameByNameRequest) (*FindBoardgameByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBoardgameByName not implemented")
}
func (UnimplementedBGGServiceServer) FindBoardgameById(context.Context, *FindBoardgameByIdRequest) (*FindBoardgameByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBoardgameById not implemented")
}
func (UnimplementedBGGServiceServer) InsertNewBoardGame(context.Context, *InsertNewBoardgameRequest) (*InsertNewBoardgameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewBoardGame not implemented")
}
func (UnimplementedBGGServiceServer) mustEmbedUnimplementedBGGServiceServer() {}

// UnsafeBGGServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BGGServiceServer will
// result in compilation errors.
type UnsafeBGGServiceServer interface {
	mustEmbedUnimplementedBGGServiceServer()
}

func RegisterBGGServiceServer(s grpc.ServiceRegistrar, srv BGGServiceServer) {
	s.RegisterService(&BGGService_ServiceDesc, srv)
}

func _BGGService_FindBoardgameByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBoardgameByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGGServiceServer).FindBoardgameByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BGGService/FindBoardgameByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGGServiceServer).FindBoardgameByName(ctx, req.(*FindBoardgameByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGGService_FindBoardgameById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBoardgameByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGGServiceServer).FindBoardgameById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BGGService/FindBoardgameById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGGServiceServer).FindBoardgameById(ctx, req.(*FindBoardgameByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGGService_InsertNewBoardGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertNewBoardgameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGGServiceServer).InsertNewBoardGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BGGService/InsertNewBoardGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGGServiceServer).InsertNewBoardGame(ctx, req.(*InsertNewBoardgameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BGGService_ServiceDesc is the grpc.ServiceDesc for BGGService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BGGService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.BGGService",
	HandlerType: (*BGGServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBoardgameByName",
			Handler:    _BGGService_FindBoardgameByName_Handler,
		},
		{
			MethodName: "FindBoardgameById",
			Handler:    _BGGService_FindBoardgameById_Handler,
		},
		{
			MethodName: "InsertNewBoardGame",
			Handler:    _BGGService_InsertNewBoardGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/bgg.proto",
}
