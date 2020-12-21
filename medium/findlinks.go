package main

import (
	"fmt"
	"os"

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
	t.Run()
	var ex links.Extractor
	fmt.Println("testsrv")
	ex = links.FileName(os.Args[1])

	if os.Args[1] == "web" {
		fmt.Println("web")
		ex = links.Url(os.Args[2])
	}

	findLinks(ex)
}
