// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type fss struct {
	co int
	ff map[string]bool
}

func main() {
	//fs := make(map[string]map[string]bool)
	//counts := make(map[string]int)
	counts := make(map[string]*fss)
	files := os.Args[1:]
	if len(files) == 0 {
		//countLines2(os.Stdin, counts, nil)
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//countLines2(f, counts, fs)
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		// if n > 1 {
		// 	v := fs[line]
		// 	var res []string
		// 	for key := range v {
		// 		res = append(res, key)
		// 	}
		// 	fmt.Printf("%d\t%s: %s\n", n, line, res)
		// }
		if n.co > 1 {
			var res []string
			for key := range n.ff {
				res = append(res, key)
			}
			fmt.Printf("%d\t%s:%s\n", n.co, line, res)
		}
	}
}

func countLines(f *os.File, counts map[string]*fss) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}

		if counts[input.Text()] == nil {
			counts[input.Text()] = &fss{ff: map[string]bool{f.Name(): true}}
		} else {
			counts[input.Text()].ff[f.Name()] = true
		}
		counts[input.Text()].co++
	}
}

func countLines2(f *os.File, counts map[string]int, fs map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}

		counts[input.Text()]++
		if fs != nil {
			if fs[input.Text()] == nil {
				ff := make(map[string]bool)
				fs[input.Text()] = ff
			}
			v := fs[input.Text()]
			v[f.Name()] = true
		}
	}
}
