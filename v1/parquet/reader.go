package parquet

import "github.com/xitongsys/parquet-go/reader"

type FileRef struct {
	reader *reader.ParquetReader
	Path string;
	IsOpen bool;
}

type ParquetReader interface {
	New() *ParquetReader
	Open() (error)
	CloseFile(file FileRef) (error)
	ReadLines(startLine int64, offset int64) ([]string, error)
}