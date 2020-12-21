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
	go t.Run(done)
	var ex links.Extractor
	fmt.Println("testsrv")
	ex = links.FileName(os.Args[1])

	if os.Args[1] == "web" {
		fmt.Println("web")
		ex = links.Url(os.Args[2])
	}
	time.Sleep(3000 * time.Millisecond)

	resp, err := http.Get("http://127.0.0.1:8000/")
	for err != nil {
		time.Sleep(250 * time.Millisecond)
		log.Printf("error:%s\n", err)
		resp, err := http.Get("http://127.0.0.1:8000/")
	}

	fmt.Printf("status:%s\n", resp.StatusCode)

	findLinks(ex)
	<-done
}
