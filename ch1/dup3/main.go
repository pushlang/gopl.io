// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]string)

	args := os.Args[1:]
	args = []string{"a", "b"}

	for _, filename := range args {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
				if _, ok := files[line]; !ok || (ok && (files[line] != filename)) {
				files[line] += filename
			}
		}

	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
		}
	}
}

//!-
