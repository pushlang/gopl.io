package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	buf := make([]byte, 8)
	var rf = &readfrom{}
	ReadHandle(buf, ReaderFunc(rf.file))
	fmt.Println(string(buf))
}

type readfrom struct {
}

type ReaderFunc func(p []byte) (n int, err error)

func (f ReaderFunc) Read(p []byte) (n int, err error) {
	return f(p)
}

func ReadHandle(p []byte, r io.Reader) {
	r.Read(p)
}

func (s *readfrom) stdin(p []byte) (n int, err error) {
	return read(p, os.Stdin)
}

func (s *readfrom) file(p []byte) (n int, err error) {
	f, err := os.Create("test.txt")
	defer f.Close()
	check(err)
	fmt.Fprint(f, "test str")
	f.Seek(0, 0)
	return read(p, f)
}
func read(p []byte, r io.Reader) (int, error) {
	return io.ReadFull(r, p)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
