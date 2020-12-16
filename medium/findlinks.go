package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/medium/links"
)

func breadthFirst(f func(item string) []string, wl links.Extractor) {
	worklist, _ := wl.Extract()
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	fn := links.FileName(os.Args[1])
	breadthFirst(crawl, fn)
}
