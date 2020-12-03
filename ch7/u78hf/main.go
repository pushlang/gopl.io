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
}

func (s *handlerfunc) foo(n int) {
	fmt.Println("handlerfunc -> foo", n)
}

type handler struct {
}

func (h *handler) ServeHTTP(n int) {
	fmt.Println("handler -> ServeHTTP", n)
}

func main() {
	var h = &handler{}
	var hf = &handlerfunc{}

	Handle(123, h)

	Handle(456, HandlerFunc(hf.foo))
}
