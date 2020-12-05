package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type HandlerFunc func(int)

type Handler interface {
	ServeHTTP(int)
}

func (f HandlerFunc) ServeHTTP(n int) {
	f(n)
}

func Handle(n int, h Handler) {
	h.ServeHTTP(n)
}

type handlerfunc struct {
	i int
}

func (s *handlerfunc) foo1(n int) {
	s.i += n
	fmt.Println("handlerfunc -> foo1", n)
}

func (s *handlerfunc) foo2(n int) {
	s.i *= n
	fmt.Println("handlerfunc -> foo2", n)
}

type handler struct {
}

func (h *handler) ServeHTTP(n int) {
	fmt.Println("handler -> ServeHTTP", n)
}

func main() {
	//var h = &handler{}
	//Handle(2, h)
	//var hf = &handlerfunc{}
	//Handle(1, HandlerFunc(hf.foo1))
	//Handle(2, HandlerFunc(hf.foo2))

	buf := make([]byte, 8)
	var rf = &readfrom{}
	ReadHandle(buf, ReaderFunc(rf.file))
	fmt.Println(string(buf))
}

type ReaderFunc func(p []byte) (n int, err error)

func (f ReaderFunc) Read(p []byte) (n int, err error) {
	return f(p)
}
func ReadHandle(p []byte, r io.Reader) {
	r.Read(p)
}

type readfrom struct {
}

func (s *readfrom) stdin(p []byte) (n int, err error) {
	return read(p, os.Stdin)
}

func (s *readfrom) file(p []byte) (n int, err error) {
	f, err := os.Create("test.txt")
	w := bufio.NewWriter(f)
	w.WriteString("test str")
	//fmt.Fprint(f, "test str")
	f.Sync()
	w.Flush()
	//f, err = os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	return read(p, f)
}
func read(p []byte, r io.Reader) (int, error) {
	defer r.(*os.File).Close()
	fmt.Printf("read: %s", string(p))
	return io.ReadFull(r, p)
}
