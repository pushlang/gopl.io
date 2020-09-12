// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//var wg sync.WaitGroup

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	depth := byte(2)
	nworkers := byte(20)
	var nWorkList byte

	type workDepth struct {
		links []string
		depth byte
	}

	type unseenDepth struct {
		links string
		depth byte
	}

	worklist := make(chan workDepth)      // lists of URLs, may have duplicates
	unseenLinks := make(chan unseenDepth) // de-duplicated URLs

	// Add command-line arguments to worklist.
	nWorkList++
	go func() { worklist <- workDepth{os.Args[1:], 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := byte(0); i < nworkers; i++ {
		go func() {
			for link := range unseenLinks {
				fmt.Print(link.depth, ":")
				foundLinks := workDepth{crawl(link.links), link.depth + 1}
				fmt.Print(".")
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		nWorkList--
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				if list.depth != depth {
					nWorkList++
					unseenLinks <- unseenDepth{link, list.depth}
				}
			}
		}
		if nWorkList == 0 {
			fmt.Println("nWorkList:", nWorkList)
			close(worklist)
		}

	}
}

//!-
