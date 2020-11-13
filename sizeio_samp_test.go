package sizeio_test

import (
	"fmt"
	"log"

	"github.com/prantlf/go-sizeio"
)

func ExampleSizeReader() {
	// Add size to a reader.
	reader := sizeio.SizeReader(nil, 456)

	fmt.Printf("Size: %d", reader.Size())
	// Output:
	// Size: 456
}

func ExampleSizeReadCloser() {
	// Add size to a reader/closer.
	reader := sizeio.SizeReadCloser(nil, 789)

	fmt.Printf("Size: %d", reader.Size())
	// Output:
	// Size: 789
}

func ExampleOpenFile() {
	// Open a file for reading and obtain its size right away.
	reader, err := sizeio.OpenFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	fmt.Printf("Size: %d", reader.Size())
	// Output:
	// Size: 17
}
