# go-sizeio

[![PkgGoDev](https://pkg.go.dev/badge/github.com/prantlf/go-sizeio)](https://pkg.go.dev/github.com/prantlf/go-sizeio)
[![Test Status](https://github.com/prantlf/go-sizeio/workflows/Test/badge.svg)](https://github.com/prantlf/go-sizeio/actions)
[![codecov](https://codecov.io/gh/prantlf/go-sizeio/branch/master/graph/badge.svg?token=XS0COPSRR7)](https://codecov.io/gh/prantlf/go-sizeio)

Extends [readers] with an additional method [Size] returning the size of the reader content. See the [documentation] for more information.

## Installation

Add this package to `go.mod` and `go.sub` in your Go project:

    go get github.com/prantlf/go-sizeio

## Usage

Use convenience methods to create readers with the extra `Size` method:

```go
import "github.com/prantlf/go-sizeio"
// get a file reader/closer with the size available right away
readerWithSize := sizeio.OpenFile("...")
// get a general reader with the size known ahead
readerWithSize := sizeio.SizeReader(otherReader, 456)
// test any reader whether it supports the Size method
_, ok := someReader.(sizeio.WithSize)
// get the size of the reader content
size := readerWithSize.Size() // returns 456 as in64
```

See the [documentation] for the full interface.

[readers]: https://golang.org/pkg/io/#Reader
[Size]: https://pkg.go.dev/github.com/prantlf/go-sizeio/#WithSize
[documentation]: https://pkg.go.dev/github.com/prantlf/go-sizeio#section-documentation
