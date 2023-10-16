package main

import (
	"errors"
	"fmt"
	"os"
)

type FileNameError struct {
	Path string
	Err  error
}

func (f *FileNameError) Error() string {
	return "main. " + f.Path + "" + f.Err.Error()
}

func main() {

	_, err := openFile("test.taxt", "root")

	var q *QueryError

	if err != nil {
		if errors.As(err, &q) {
			fmt.Println("custom error found in the chain", q.Func)
			return
		}
		return
	}

}

func openFile(fileName string, dirName string) (*os.File, error) {

	f, err := os.Open((fileName))

	if err != nil {
		return nil, &FileNameError{

			Path: dirName,
			Err:  err,
		}
	}
	return f, nil
}
