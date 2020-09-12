package main

import "fmt"

func main() {
	tb := tB{"tB"}
	tc := tC{"tC"}
	foo(tb)
	foo(tc)
}

type iA interface{ A() }

type tB struct{ s string }

type tC struct{ s string }

func (tb tB) A() { fmt.Println(tb.s) }

func (tc tC) A() { fmt.Println(tc.s) }

func foo(ia iA) { ia.A() }

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
