package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/coffeeandcloud/douglasie-server/v1/parquet"
	"github.com/coffeeandcloud/douglasie-server/v1/rpc"
	grpc "google.golang.org/grpc"
)

type parquetServer struct {
	ref *parquet.FileRef
	rpc.UnimplementedParquetServer
}

func (s *parquetServer) OpenFile(ctx context.Context, in *rpc.OpenFileReq) (*rpc.FileInfoResp, error) {
	s.ref = &parquet.FileRef{
		Path: in.Filename,
	}
	err := s.ref.Open()
	if err != nil {
		return nil, err
	}
	numOfRows, _ := s.ref.GetNumRows()
	return &rpc.FileInfoResp{
		Filename:  s.ref.Path,
		NumOfRows: numOfRows,
	}, nil
}

func (s *parquetServer) ReadRows(in *rpc.GetRowsReq, stream rpc.Parquet_ReadRowsServer) error {
	rows, err := s.ref.ReadLines(int64(in.StartLine), int64(in.Offset))
	if err != nil {
		return err
	}
	for _, row := range rows {
		for k, v := range row {
			stream.SendMsg(&rpc.Row{
				Key: k,
				Val: fmt.Sprintf("%s", v),
			})
		}
	}

	return nil
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
