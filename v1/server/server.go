package server

import (
	"log"
	"net"

	"github.com/coffeeandcloud/douglasie-server/v1/rpc"
	grpc "google.golang.org/grpc"
)

type parquetServer struct {
	rpc.UnimplementedParquetServer
}

func Run() {
	lis, err := net.Listen("unix", "/tmp/douglasie.sock")

	if err != nil {
		log.Panicln(err)
	}

	s := grpc.NewServer()
	rpc.RegisterParquetServer(s, &parquetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
