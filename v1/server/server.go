package server

import (
	"context"
	"log"
	"net"

	"github.com/coffeeandcloud/douglasie-server/v1/rpc"
	grpc "google.golang.org/grpc"
)

type parquetServer struct {
	rpc.UnimplementedParquetServer
}

func (s *parquetServer) OpenFile(ctx context.Context, in *rpc.OpenFileReq) (*rpc.FileInfoResp, error) {
	return &rpc.FileInfoResp{
		Filename:  in.Filename,
		NumOfRows: 987,
	}, nil
}

func Run() {
	// lis, err := net.Listen("unix", "/tmp/douglasie.sock")
	lis, err := net.Listen("tcp", ":5555")

	if err != nil {
		log.Panicln(err)
	}

	s := grpc.NewServer()
	rpc.RegisterParquetServer(s, &parquetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
