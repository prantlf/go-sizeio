// Package sizeio extends interfaces Reader and ReadCloser with the interface
// WithSize, which provides the size of the reader content.
//
// Strings and bytes support the Size method out-of-tbe-box:
//
//     readerWithSize := strings.NewReader("...")
//     readerWithSize := bytes.NewReader([]byte{...})
//
// Files can be opened or converted using the following convenience methods:
//
//     readerWithSize := sizeio.OpenFile("...")
//     readerWithSize := sizeio.SizeFile(file)
//
// Readers can be converted using the following convenience methods:
//
//     readerWithSize := sizeio.SizeReader(reader, 456)
//     readerWithSize := sizeio.SizeReadCloser(readClose, 789)
//
// Otherwise you can use a structure with the following method, which
// is essentially the WithSize interface:
//
//     Size() int64
//
// And finally, test a reader and get the size of the reader content:
//
//     _, ok := reader.(sizeio.WithSize) // returns true if Size is provided
//     size := readerWithSize.Size()     // returns 789 as in64
package sizeio

import (
	"io"
	"os"
)

// WithSize provides the size of the reader content. Used by
// ReaderWithSize and ReadCloserWithSize.
type WithSize interface {
	// Size returns the size of the content delivered by a reader.
	Size() int64
}

// ReaderWithSize encapsulates two interfaces - Reader and WithSize.
// It is returned by SizeReader.
type ReaderWithSize interface {
	io.Reader
	WithSize
}

type readerWithSize struct {
	io.Reader
	size int64
}

func (r *readerWithSize) Size() int64 {
	return r.size
}

// SizeReader adds the interface WithSize to a Reader instance.
func SizeReader(reader io.Reader, size int64) ReaderWithSize {
	return &readerWithSize{reader, size}
}

// ReadCloserWithSize encapsulates two interfaces - ReadCloser and WithSize.
// It is returned by SizeReadCloser.
type ReadCloserWithSize interface {
	io.ReadCloser
	WithSize
}

type readCloserWithSize struct {
	io.ReadCloser
	size int64
}

func (r *readCloserWithSize) Size() int64 {
	return r.size
}

// SizeReadCloser adds the interface WithSize to a ReadCloser instance.
func SizeReadCloser(reader io.ReadCloser, size int64) ReadCloserWithSize {
	return &readCloserWithSize{reader, size}
}

// SizeFile gets the file size from an opened file and adds to the resulting
// ReaderCloser within the interface WithSize.
//
// The input File will be wrapped in the output ReaderCloser. Do not forget
// to close it, once you do not need it, or defer the closure to perform
// it automatically in case of a failure.
func SizeFile(file *os.File) (ReadCloserWithSize, error) {
	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}
	return &readCloserWithSize{file, stat.Size()}, nil
}

// OpenFile opens the named file for reading only. It obtains the file size
// right away to support the WithSize interface.
//
// Do not forget to close the reader, once you do not need it, or defer
// the closure to perform it automatically in case of a failure.
func OpenFile(filePath string) (ReadCloserWithSize, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return SizeFile(file)
}
