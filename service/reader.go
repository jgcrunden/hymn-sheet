package service

import "os"

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type OSFileReader struct{}

func NewOSFileReader() OSFileReader {
	return OSFileReader{}
}

func (o OSFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
