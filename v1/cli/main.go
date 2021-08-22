package main

import (
	"context"
	"log"
	"time"

	"github.com/coffeeandcloud/douglasie-server/v1/rpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewParquetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.OpenFile(ctx, &rpc.OpenFileReq{
		Filename: "sample/flat.parquet",
	})
	if err != nil {
		log.Fatalf("could not open parquet: %v", err)
	}
	log.Printf("File '%s' has %d rows.", r.Filename, r.NumOfRows)
}
