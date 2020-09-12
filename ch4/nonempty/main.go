// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//func(d []string) {data=d}(nonempty2(data))
//!-nonempty

func main() {
	//!+main
	//data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", func(d []string) []string { data = d; return data }(nonempty2(data))) // `["one" "three"]`
	//fmt.Printf("%q\n", data)                                                                 // `["one" "three" "three"]`
	data2 := []string{"a", "b", "b", "c", "c", "c"}
	fmt.Printf("%q\n", dup(data2))

	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func dup(strings []string) []string {
	i := 1
	l := 0

	for j := 0; j < len(strings)-1-l; j++ {
		for k := j + 1; k < len(strings)-l; k++ {
			if strings[j] != strings[k] {
				strings[i] = strings[k]
				i++
				break
			}
			l++
		}
		fmt.Println(i, j, l)
	}

	return strings[:i]
}

//!-alt
