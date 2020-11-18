package sizeio_test

import (
	"os"
	"testing"

	"github.com/prantlf/go-sizeio"
)

func TestOpenFile_close(t *testing.T) {
	reader, err := sizeio.OpenFile("demo/test.txt")
	if err != nil {
		t.Error("sizeio: valid file failed -", err)
	}
	if reader == nil {
		t.Error("sizeio: valid file nil")
	}
	if err := reader.Close(); err != nil {
		t.Error("sizeio: closing file failed -", err)
	}
}

func TestOpenFile_missing(t *testing.T) {
	reader, err := sizeio.OpenFile("missing.txt")
	if err == nil {
		t.Error("sizeio: invalid file opened")
	}
	if reader != nil {
		t.Error("sizeio: invalid reader not nil")
	}
}

func TestOpenFile_closed(t *testing.T) {
	file, err := os.Open("demo/test.txt")
	if err != nil {
		t.Error("sizeio: valid file failed -", err)
	}
	file.Close()
	_, err = sizeio.SizeFile(file)
	if err == nil {
		t.Error("sizeio: closed file inspected")
	}
}
