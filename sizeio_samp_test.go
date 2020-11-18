package sizeio_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/prantlf/go-sizeio"
	"github.com/prantlf/go-sizeio/demo"
)

func Example() {
	// Open a file for reading and obtain its size right away.
	reader, err := sizeio.OpenFile("demo/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// Test any reader whether it supports the Size method
	_, ok := reader.(sizeio.WithSize)

	// Get the size of the content delivered by the reader.
	size := reader.Size()

	fmt.Printf("Supported: %v\n", ok)
	fmt.Printf("Size:      %d\n", size)
	fmt.Printf("Content:   %s\n", demo.StringifyReader(reader))
	// Output:
	// Supported: true
	// Size:      17
	// Content:   text file content
}

func ExampleWithSize() {
	input := strings.NewReader("content")

	// Add size to a reader.
	reader := sizeio.SizeReader(input, 7)

	_, ok := reader.(sizeio.WithSize)
	fmt.Printf("Supported: %v\n", ok)
	// Output:
	// Supported: true
}

func ExampleWithSize_Size() {
	input := strings.NewReader("content")

	// Add size to a reader.
	reader := sizeio.SizeReader(input, 7)

	if readerWithSize, ok := reader.(sizeio.WithSize); ok {
		fmt.Printf("Size: %v\n", readerWithSize.Size())
	}
	// Output:
	// Size: 7
}

func ExampleSizeReader() {
	input := strings.NewReader("content")

	// Add size to a reader.
	reader := sizeio.SizeReader(input, 7)

	fmt.Printf("Size: %d\n", reader.Size())
	// Output:
	// Size: 7
}

func ExampleSizeReadCloser() {
	input := ioutil.NopCloser(strings.NewReader("content"))

	// Add size to a reader/closer.
	reader := sizeio.SizeReadCloser(input, 7)
	defer reader.Close()

	fmt.Printf("Size: %d\n", reader.Size())
	// Output:
	// Size: 7
}

func ExampleSizeFile() {
	file, err := os.Open("demo/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Add the file size to the file reader.
	reader, err := sizeio.SizeFile(file)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	defer reader.Close()

	fmt.Printf("Size: %d\n", reader.Size())
	// Output:
	// Size: 17
}

func ExampleOpenFile() {
	// Open a file for reading and obtain its size right away.
	reader, err := sizeio.OpenFile("demo/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	fmt.Printf("Size: %d\n", reader.Size())
	// Output:
	// Size: 17
}
