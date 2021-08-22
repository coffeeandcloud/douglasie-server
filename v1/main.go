package main

import (
	"log"

	"github.com/coffeeandcloud/douglasie-server/v1/server"
)

func main() {
	log.Println("Welcome to Douglasie Parquet Reader!")

	/*
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
	*/

	server.Run()
}
