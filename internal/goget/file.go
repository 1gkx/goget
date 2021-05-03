package goget

import (
	"io"
	"os"
)

type File struct {
	fileName string
	f        *os.File
}

func CreateFile(filepath string) *File {

	file, _ := os.Create(filepath + ".tmp")
	return &File{
		fileName: filepath,
		f:        file,
	}

}

func (file *File) Close() {
	file.f.Close()
}

func (file *File) Save(body io.Reader) error {

	counter := &WriteCounter{}
	_, err := io.Copy(file.f, io.TeeReader(body, counter))

	err = os.Rename(file.fileName+".tmp", file.fileName)
	if err != nil {
		return err
	}

	return err

}
