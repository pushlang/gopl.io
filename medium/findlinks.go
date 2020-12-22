package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopl.io/medium/links"
	t "gopl.io/medium/testsrv"
)

func findLinks(wl links.Extractor) {
	worklist, _ := wl.Extract()

	for _, w := range worklist {
		fmt.Println(w)
	}
}

func main() {
	done := make(chan struct{})
	var ex links.Extractor
	ex = links.FileName(os.Args[1])

	if os.Args[1] == "web" {
		fmt.Println("web")
		ex = links.Url(os.Args[2])
	}

	go func() {
		req, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)

		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)

		for err != nil {
			time.Sleep(250 * time.Millisecond)
			log.Printf("error:%s\n", err)
			resp, err = http.DefaultClient.Do(req)
		}
		fmt.Printf("status code:%d\n", resp.StatusCode)

		findLinks(ex)
	}()

	t.Run()
}
