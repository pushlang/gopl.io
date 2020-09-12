// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := [...]int{0, 1, 2, 3, 4, 5}
	reverse2(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	fmt.Println("b", b, "len: ", len(b), "cap: ", cap(b))
	b = rotate2(b, 3)
	fmt.Println("b", b, "len: ", len(b), "cap: ", cap(b))

	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) []int {
	if n >= len(s) || n <= 0 {
		return []int{}
	}

	return append(s, s[:n]...)[n:]
}

func rotate2(s [6]int, n int) [6]int {
	for i := 0; i < n; i++ {
		for j := 0; j < len(s)-1; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}

	return s
}

//!-rev
