package main

import (
	"log"

	"github.com/coffeeandcloud/douglasie-server/v1/parquet"
)

func main() {
	log.Println("Welcome to Douglasie Parquet Reader!")
	ref := parquet.FileRef{
		Path: "sample/flat.parquet",
	}
	err := ref.Open()
	if err != nil {
		log.Println(err)
	}
}
