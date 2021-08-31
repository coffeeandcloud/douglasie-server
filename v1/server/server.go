package server

import (
	"log"

	"github.com/coffeeandcloud/douglasie-server/v1/parquet"
	"github.com/gin-gonic/gin"
)

func Run() {

	r := gin.Default()

	r.GET("/info/*action", func(c *gin.Context) {
		path := c.Params.ByName("action")
		log.Printf("URL: %s", path)

		ref := &parquet.FileRef{
			Path: path,
		}
		err := ref.Open()

		if err != nil {
			log.Fatalf("Error opening file '%s': %s", ref.Path, err)
		}

		numOfRows, _ := ref.GetNumRows()

		c.JSON(200, gin.H{
			"numOfRows": numOfRows,
			"Path":      ref.Path,
		})

		ref.CloseFile()
	})

	r.GET("/content/*action", func(c *gin.Context) {
		startLine := 0
		offset := 10

		path := c.Params.ByName("action")

		ref := &parquet.FileRef{
			Path: path,
		}
		err := ref.Open()

		rows, err := ref.ReadLines(int64(startLine), int64(offset))
		if err != nil {
			log.Fatalf("Error opening file '%s': %s", ref.Path, err)
		}
		c.JSON(200, gin.H{
			"items": rows,
		})
	})

	socket := "/tmp/douglasie.sock"
	err := r.RunUnix(socket)

	if err != nil {
		log.Fatalf("An error occurred on the socket '%s': %s", socket, err)
	}
}
