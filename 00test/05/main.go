package main

import "fmt"

var tb = &tB{}

func main() {
	foo(fA(tb.fB))
	fmt.Println(tb.s)
	foo(fA(tb.fC))
	fmt.Println(tb.s)

}

type iA interface {
	A()
}

type fA func()

func (fa fA) A() {
	fa()
}

type tB struct{ s string }

func (tb *tB) fB() {
	tb.s = "tB: fB"
	//fmt.Println("tb.fB")
}

func (tb *tB) fC() {
	tb.s = "tB: fC"
	//fmt.Println("tb.fC")
}

func foo(ia iA) {
	ia.A()
}

// mux := http.NewServeMux() //func http.NewServeMux() *http.ServeMux
// mux.Handle("/list", http.HandlerFunc(db.list))

// func (mux *ServeMux) Handle(pattern string, handler Handler) {
// }

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// type HandlerFunc func(ResponseWriter, *Request)

// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }

// func (db database) list(w http.ResponseWriter, req *http.Request) {
// }

// type database map[string]dollars
