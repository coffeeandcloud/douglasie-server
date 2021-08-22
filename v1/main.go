package main

import (
	"log"

	"github.com/coffeeandcloud/douglasie-server/v1/parquet"
	"github.com/coffeeandcloud/douglasie-server/v1/server"
)

func main() {
	log.Println("Welcome to Douglasie Parquet Reader!")

	ref := parquet.FileRef{
		Path: "sample/flat.parquet",
	}
	err := ref.Open()
	if err != nil {
		log.Panicln(err)
	}

	rows, err := ref.ReadLines(0, 1000)
	if err != nil {
		log.Panicln(err)
	}

	schema, _ := ref.GetSchema()
	log.Println(schema.Children)

	for i := 0; i < len(rows); i++ {
		for k, v := range rows[i] {
			log.Printf("Row %d: ", i)
			log.Printf("Key: %s Val: %s", k, v)
		}
	}

	server.Run()
}
