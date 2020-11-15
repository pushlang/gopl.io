//rm u78; goimports -v -w main.go ./count/count.go; gofmt -w main.go ./count/count.go; go build; ./u78
package main

import (
	"log"
	"net/http"

	h "gopl.io/ch7/u78/html"
)

func main() {
	http.HandleFunc("/sort", h.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
