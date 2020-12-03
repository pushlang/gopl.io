package main

import "fmt"

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
	s.i*=n
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

	var hf = &handlerfunc{}

	Handle(1, HandlerFunc(hf.foo1))
	Handle(2, HandlerFunc(hf.foo2))
}

type ReaderFunc(p []byte) (n int, err error)

func (f ReaderFunc) ReadSomthing(p []byte) (n int, err error) {
	return f(p)
}

func ReadHandle(p []byte) (n int, err error) {
	h.ReadSomthing(n)
}

type readfrom struct {
	s string
}

func (s *readfrom) stdin(p []byte) (n int, err error) {
	s.i += n
	fmt.Println("handlerfunc -> foo1", n)
}

func (s *readfrom) strng(p []byte) (n int, err error) {
	s.s=string(p)
	return len(
}