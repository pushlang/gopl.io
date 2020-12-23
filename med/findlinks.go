package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/med/links"
	"gopl.io/med/testsrv"
)

func findLinks(wl links.Extractor, done chan struct{}) {
	log.Println("go findLinks()")
	worklist, _ := wl.Extract()

	for _, w := range worklist {
		fmt.Println(w)
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	if len(os.Args) == 1 {
		log.Fatal("No arguments")
	}

	var ex links.Extractor

	if os.Args[1] == "web" {
		log.Println("Web extractor")
		if len(os.Args) <= 2 {
			ex = links.Url("http://127.0.0.1:8000")
			log.Println("Using default http://127.0.0.1:8000")
		} else {
			ex = links.Url(os.Args[2])
			log.Println("using " + os.Args[2])
		}
	} else if os.Args[1] == "file" {
		log.Println("File extractor")
		ex = links.FileName(os.Args[2])
	} else {
		log.Fatal("Unknown action")
	}
	go findLinks(ex, done)

	if os.Args[1] == "web" && len(os.Args) <= 2 {
		testsrv.Run()
	}
	<-done
}
