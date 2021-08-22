package parquet

import (
	"log"
	"reflect"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/schema"
)

func (ref *FileRef) Open() error {
	file, err := local.NewLocalFileReader(ref.Path)
	if err != nil {
		log.Println("Can't open file 🚨")
		return err
	}

	ref.reader, err = reader.NewParquetReader(file, nil, 4)

	if err != nil {
		log.Println("Can't create file reader 🙈")
		return err
	} else {
		log.Println("NumRows:", ref.reader.GetNumRows())
	}
	return nil
}

func (ref *FileRef) ReadLines(startLine int64, offset int64) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)
	ref.reader.SkipRows(startLine)
	res, _ := ref.reader.ReadByNumber(int(offset))

	for i := 0; i < len(res); i++ {
		row := make(map[string]interface{}, 1)
		numFields := reflect.ValueOf(res[i]).NumField()
		for j := 0; j < numFields; j++ {
			fieldName := reflect.TypeOf(res[i]).Field(j).Name
			row[fieldName] = reflect.ValueOf(res[i]).Field(j).Interface()
		}
		result = append(result, row)
	}
	return result, nil
}

func (ref *FileRef) GetSchema() (*schema.PathMapType, error) {
	return ref.reader.SchemaHandler.PathMap, nil
}
