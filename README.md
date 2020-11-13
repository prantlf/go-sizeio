# go-sizeio

Extends [readers] with an additional method [Size] returning the size of the reader content. See the [documentation] for more information.

## Installation

Add this package to `go.mod` and `go.sub` in your Go project:

    go get github.com/prantlf/go-sizeio

## Usage

Use convenience methods to create readers with the extra `Size` method:

```go
import (
  "github.com/prantlf/go-sizeio"
)
// get a file reader with the size available right away
readerWithSize := sizeio.OpenFile("...")
// get a general reader with the size known ahead
readerWithSize := sizeio.SizeReader(otherReader, 456)
// get the size of the reader content
size := readerWithSize.Size() // returns 456 as in64
```

See the [documentation] for the full interface.

[readers]: https://golang.org/pkg/io/#Reader
[Size]: https://pkg.go.dev/github.com/prantlf/go-sizeio/#WithSize
[documentation]: https://pkg.go.dev/github.com/prantlf/go-sizeio
