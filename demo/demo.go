package demo

import (
	"io"
	"log"
	"strings"
)

func StringifyReader(reqBody io.Reader) string {
	bld := new(strings.Builder)
	if _, err := io.Copy(bld, reqBody); err != nil {
		log.Fatal(err)
	}
	return bld.String()
}
