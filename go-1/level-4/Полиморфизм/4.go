package main

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	 Read() []byte
}

type ReaderWriter interface {
	Writer
	Reader
}

type UpperReaderWriter struct {
	UpperString string
}

func (u *UpperReaderWriter) Read() []byte {
	return []byte(u.UpperString)
}

func (u *UpperReaderWriter) Write(p []byte) int {
	u.UpperString = strings.ToUpper(string(p))
	return len(p)
}

func main() {
	var ex ReaderWriter = &UpperReaderWriter{
		UpperString: "Hello",
	}
	fmt.Println(ex.Read())
	ex.Write([]byte{208, 144, 208, 145, 208, 146})
	fmt.Println(ex.Read())
}
