// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

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

// ParquetClient is the client API for Parquet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParquetClient interface {
	OpenFile(ctx context.Context, in *OpenFileReq, opts ...grpc.CallOption) (*FileInfoResp, error)
	ReadRows(ctx context.Context, in *GetRowsReq, opts ...grpc.CallOption) (Parquet_ReadRowsClient, error)
}

type parquetClient struct {
	cc grpc.ClientConnInterface
}

func NewParquetClient(cc grpc.ClientConnInterface) ParquetClient {
	return &parquetClient{cc}
}

func (c *parquetClient) OpenFile(ctx context.Context, in *OpenFileReq, opts ...grpc.CallOption) (*FileInfoResp, error) {
	out := new(FileInfoResp)
	err := c.cc.Invoke(ctx, "/rpc.Parquet/OpenFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parquetClient) ReadRows(ctx context.Context, in *GetRowsReq, opts ...grpc.CallOption) (Parquet_ReadRowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Parquet_ServiceDesc.Streams[0], "/rpc.Parquet/ReadRows", opts...)
	if err != nil {
		return nil, err
	}
	x := &parquetReadRowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Parquet_ReadRowsClient interface {
	Recv() (*Row, error)
	grpc.ClientStream
}

type parquetReadRowsClient struct {
	grpc.ClientStream
}

func (x *parquetReadRowsClient) Recv() (*Row, error) {
	m := new(Row)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ParquetServer is the server API for Parquet service.
// All implementations must embed UnimplementedParquetServer
// for forward compatibility
type ParquetServer interface {
	OpenFile(context.Context, *OpenFileReq) (*FileInfoResp, error)
	ReadRows(*GetRowsReq, Parquet_ReadRowsServer) error
	mustEmbedUnimplementedParquetServer()
}

// UnimplementedParquetServer must be embedded to have forward compatible implementations.
type UnimplementedParquetServer struct {
}

func (UnimplementedParquetServer) OpenFile(context.Context, *OpenFileReq) (*FileInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenFile not implemented")
}
func (UnimplementedParquetServer) ReadRows(*GetRowsReq, Parquet_ReadRowsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadRows not implemented")
}
func (UnimplementedParquetServer) mustEmbedUnimplementedParquetServer() {}

// UnsafeParquetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParquetServer will
// result in compilation errors.
type UnsafeParquetServer interface {
	mustEmbedUnimplementedParquetServer()
}

func RegisterParquetServer(s grpc.ServiceRegistrar, srv ParquetServer) {
	s.RegisterService(&Parquet_ServiceDesc, srv)
}

func _Parquet_OpenFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParquetServer).OpenFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Parquet/OpenFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParquetServer).OpenFile(ctx, req.(*OpenFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parquet_ReadRows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRowsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ParquetServer).ReadRows(m, &parquetReadRowsServer{stream})
}

type Parquet_ReadRowsServer interface {
	Send(*Row) error
	grpc.ServerStream
}

type parquetReadRowsServer struct {
	grpc.ServerStream
}

func (x *parquetReadRowsServer) Send(m *Row) error {
	return x.ServerStream.SendMsg(m)
}

// Parquet_ServiceDesc is the grpc.ServiceDesc for Parquet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parquet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Parquet",
	HandlerType: (*ParquetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenFile",
			Handler:    _Parquet_OpenFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRows",
			Handler:       _Parquet_ReadRows_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/rpc/parquet.proto",
}
