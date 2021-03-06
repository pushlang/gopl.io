package testsrv

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	log.Println("enter Run (test server)")
	defer log.Println("exit Run (test server)")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var links = map[string][]string{
	"/":    {"/1", "/2", "/3", "/4", "/5"},
	"/1":   {"/11", "/12", "/13", "/14", "/15"},
	"/2":   {"/21", "/22", "/23", "/24", "/25"},
	"/3":   {"/31", "/32", "/33", "/34", "/35"},
	"/4":   {"/41", "/42", "/43", "/44", "/45"},
	"/5":   {"/51", "/52", "/53", "/54", "/55"},
	"/11":  {"/111", "/112"},
	"/12":  {"/121", "/122", "/123"},
	"/23":  {"/231"},
	"/123": {"/1231", "/1232"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request path:", r.URL.Path)

	host := "127.0.0.1:8000"

	for _, link := range links[r.URL.Path] {
		fmt.Fprintf(w, "<a href=http://%s%s>%s</a><br>\n", host, link, link)
	}
}
