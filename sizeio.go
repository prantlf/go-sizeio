// Package sizeio extends interfaces Reader and ReadCloser with the interface
// WithSize, which provides the size of the reader content.
//
// Strings and bytes support the Size method out-of-tbe-box:
//
//     readerWithSize := strings.NewReader("...")
//     readerWithSize := bytes.NewReader([]byte{...})
//
// Files can be opened using a convenience method to get a special reader
// with the size included:
//
//     readerWithSize := sizeio.OpenFile("...")
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
// And finally, get the size of the reader content:
//
//     size := readerWithSize.Size() // returns 789 as in64
package sizeio

import (
	"io"
	"os"
)

// WithSize provides the size of the reader content.
type WithSize interface {
	// Size returns the size of the reader content.
	Size() int64
}

// ReaderWithSize encapsulates two interfaces Reader and WithSize.
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

// SizeReader adds the interface WithSize to an Reader instance.
func SizeReader(reader io.Reader, size int64) ReaderWithSize {
	return &readerWithSize{reader, size}
}

// ReadCloserWithSize encapsulates two interfaces ReadCloser and WithSize.
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

// SizeReadCloser adds the interface WithSize to an ReadCloser instance.
func SizeReadCloser(reader io.ReadCloser, size int64) ReadCloserWithSize {
	return &readCloserWithSize{reader, size}
}

// OpenFile opens the named file for reading only. It obtains the file size
// right away to support the WithSize interface.
func OpenFile(filePath string) (ReadCloserWithSize, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}
	return SizeReadCloser(file, stat.Size()), nil
}
