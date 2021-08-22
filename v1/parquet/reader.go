package parquet

import (
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/schema"
)

type FileRef struct {
	reader    *reader.ParquetReader
	Path      string
	IsOpen    bool
	NumOfRows int64
}

type ParquetReader interface {
	New() *ParquetReader
	Open() error
	CloseFile() error
	ReadLines(startLine int64, offset int64) ([]interface{}, error)
	GetSchema() (*schema.PathMapType, error)
	GetNumRows() (int64, error)
}
