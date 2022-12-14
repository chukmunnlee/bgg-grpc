// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: grpc/server/bgg.proto

package server

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
	FindCommentsByGameId(ctx context.Context, in *FindCommentsByGameIdRequest, opts ...grpc.CallOption) (BGGService_FindCommentsByGameIdClient, error)
	InsertNewComment(ctx context.Context, in *InsertNewCommentRequest, opts ...grpc.CallOption) (*InsertNewCommentResponse, error)
}

type bGGServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBGGServiceClient(cc grpc.ClientConnInterface) BGGServiceClient {
	return &bGGServiceClient{cc}
}

func (c *bGGServiceClient) FindBoardgameByName(ctx context.Context, in *FindBoardgameByNameRequest, opts ...grpc.CallOption) (*FindBoardgameByNameResponse, error) {
	out := new(FindBoardgameByNameResponse)
	err := c.cc.Invoke(ctx, "/grpc.server.BGGService/FindBoardgameByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGGServiceClient) FindBoardgameById(ctx context.Context, in *FindBoardgameByIdRequest, opts ...grpc.CallOption) (*FindBoardgameByIdResponse, error) {
	out := new(FindBoardgameByIdResponse)
	err := c.cc.Invoke(ctx, "/grpc.server.BGGService/FindBoardgameById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGGServiceClient) InsertNewBoardGame(ctx context.Context, in *InsertNewBoardgameRequest, opts ...grpc.CallOption) (*InsertNewBoardgameResponse, error) {
	out := new(InsertNewBoardgameResponse)
	err := c.cc.Invoke(ctx, "/grpc.server.BGGService/InsertNewBoardGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bGGServiceClient) FindCommentsByGameId(ctx context.Context, in *FindCommentsByGameIdRequest, opts ...grpc.CallOption) (BGGService_FindCommentsByGameIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &BGGService_ServiceDesc.Streams[0], "/grpc.server.BGGService/FindCommentsByGameId", opts...)
	if err != nil {
		return nil, err
	}
	x := &bGGServiceFindCommentsByGameIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BGGService_FindCommentsByGameIdClient interface {
	Recv() (*FindCommentsByGameIdResponse, error)
	grpc.ClientStream
}

type bGGServiceFindCommentsByGameIdClient struct {
	grpc.ClientStream
}

func (x *bGGServiceFindCommentsByGameIdClient) Recv() (*FindCommentsByGameIdResponse, error) {
	m := new(FindCommentsByGameIdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bGGServiceClient) InsertNewComment(ctx context.Context, in *InsertNewCommentRequest, opts ...grpc.CallOption) (*InsertNewCommentResponse, error) {
	out := new(InsertNewCommentResponse)
	err := c.cc.Invoke(ctx, "/grpc.server.BGGService/InsertNewComment", in, out, opts...)
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
	FindCommentsByGameId(*FindCommentsByGameIdRequest, BGGService_FindCommentsByGameIdServer) error
	InsertNewComment(context.Context, *InsertNewCommentRequest) (*InsertNewCommentResponse, error)
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
func (UnimplementedBGGServiceServer) FindCommentsByGameId(*FindCommentsByGameIdRequest, BGGService_FindCommentsByGameIdServer) error {
	return status.Errorf(codes.Unimplemented, "method FindCommentsByGameId not implemented")
}
func (UnimplementedBGGServiceServer) InsertNewComment(context.Context, *InsertNewCommentRequest) (*InsertNewCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewComment not implemented")
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
		FullMethod: "/grpc.server.BGGService/FindBoardgameByName",
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
		FullMethod: "/grpc.server.BGGService/FindBoardgameById",
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
		FullMethod: "/grpc.server.BGGService/InsertNewBoardGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGGServiceServer).InsertNewBoardGame(ctx, req.(*InsertNewBoardgameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BGGService_FindCommentsByGameId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindCommentsByGameIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BGGServiceServer).FindCommentsByGameId(m, &bGGServiceFindCommentsByGameIdServer{stream})
}

type BGGService_FindCommentsByGameIdServer interface {
	Send(*FindCommentsByGameIdResponse) error
	grpc.ServerStream
}

type bGGServiceFindCommentsByGameIdServer struct {
	grpc.ServerStream
}

func (x *bGGServiceFindCommentsByGameIdServer) Send(m *FindCommentsByGameIdResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BGGService_InsertNewComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertNewCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGGServiceServer).InsertNewComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.server.BGGService/InsertNewComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGGServiceServer).InsertNewComment(ctx, req.(*InsertNewCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BGGService_ServiceDesc is the grpc.ServiceDesc for BGGService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BGGService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.server.BGGService",
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
		{
			MethodName: "InsertNewComment",
			Handler:    _BGGService_InsertNewComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindCommentsByGameId",
			Handler:       _BGGService_FindCommentsByGameId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/server/bgg.proto",
}
