package sizeio_test

import (
	"testing"

	"github.com/prantlf/go-sizeio"
)

func TestOpenFile_close(t *testing.T) {
	reader, err := sizeio.OpenFile("test.txt")
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
