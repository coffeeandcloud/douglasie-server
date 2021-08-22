package parquet

import (
	"log"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func (fileRef *FileRef) Open() error {
	file, err := local.NewLocalFileReader(fileRef.Path)
	if err != nil {
		log.Println("Can't open file 🚨")
		return err
	}

	fileRef.reader, err = reader.NewParquetReader(file, nil, 4)

	if err != nil {
		log.Println("Can't create file reader 🙈")
		return err
	} else {
		log.Println("NumRows:", fileRef.reader.GetNumRows())
	}
	return nil
}
