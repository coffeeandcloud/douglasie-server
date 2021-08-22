package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/coffeeandcloud/douglasie-server/v1/rpc"
	"google.golang.org/grpc"
)

func main() {

	if len(os.Args) != 4 {
		log.Fatalln("Missing params. Use like main.go <filepath> <startline> <offset>")
	}

	path := os.Args[1]
	line, _ := strconv.ParseInt(os.Args[2], 10, 32)
	offset, _ := strconv.ParseInt(os.Args[3], 10, 32)

	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewParquetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.OpenFile(ctx, &rpc.OpenFileReq{
		Filename: path,
	})
	if err != nil {
		log.Fatalf("could not open parquet: %v", err)
	}
	log.Printf("File '%s' has %d rows.", r.Filename, r.NumOfRows)

	rowsClient, err := c.ReadRows(ctx, &rpc.GetRowsReq{
		Path:      "sample/flat.parquet",
		StartLine: int32(line),
		Offset:    int32(offset),
	})
	if err != nil {
		log.Fatalf("could not open parquet: %v", err)
	}

	//rows := make([][]string, 0)
	header := make([]string, 0)

	for {
		row, err := rowsClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", rowsClient, err)
		}

		log.Println(row)

		data := make(map[string]interface{}, 0)
		err = json.Unmarshal(row.Fields, &data)
		if err != nil {
			log.Fatalln(err)
		}

		cpRow := make([]string, 0)
		header = make([]string, 0)
		for k, v := range data {
			cpRow = append(cpRow, ToString(v))
			header = append(header, k)
		}
		//rows = append(rows, cpRow)
		//log.Println(cpRow)
	}

	/*
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		for _, v := range rows {
			table.Append(v)
		}
		table.Render()
	*/
}

func ToString(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case int32:
		return fmt.Sprintf("%d", value)
	case int64:
		return fmt.Sprintf("%d", value)
	case string:
		return fmt.Sprintf("%s", value)
	case bool:
		return fmt.Sprintf("%b", value)
	case float64:
		return fmt.Sprintf("%f", value)
	}
	return ""
}
