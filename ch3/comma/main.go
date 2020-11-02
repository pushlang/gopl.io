// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	os.Args = []string{"", "10", "100", "1000", "10000", "100000", "1000000", "10000000", "100000000"}
	//os.Args = []string{"", "-10,00", "-100,00", "-1000,00", "-10000,00", "-100000,00", "-1000000,00", "-10000000,00", "-100000000,00"}

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func commaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	fp := n % 3

	if fp == 0 {
		fp = 3
	}

	var buf bytes.Buffer

	buf.WriteString(s[:fp])

	for i := fp; i < n; i += 3 {
		buf.WriteString(".")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func comma2(s string) string {
	var buf bytes.Buffer

	trail := ""
	sign := ""

	if s[0] == '-' {
		sign = "-"
		s = s[1:]
	}

	commaIndex := bytes.LastIndex([]byte(s), []byte{','})
	if commaIndex != -1 {
		trail = s[commaIndex:]
		s = s[:commaIndex]
	}

	for i, l := 0, len(s)-1; l >= 0; l, i = l-1, i+1 {
		buf.WriteByte(s[i])

		if ((l)%3) == 0 && l > 2 {
			buf.WriteByte('.')
		}
	}

	return sign + buf.String() + trail
}

func ana(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v1 := range s1 {
		i1, i2 := 0, 0
		for _, v12 := range s1 {
			if v1 == v12 {
				i1++
			}
		}
		for _, v2 := range s2 {
			if v1 == v2 {
				i2++
			}
		}

		if i1 != i2 {
			return false
		}
	}

	return true
}

func ana2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	v12 := make(map[rune]int)
	v22 := make(map[rune]int)

	for _, v1 := range s1 {
		v12[v1]++
	}

	for _, v2 := range s2 {
		v22[v2]++
	}

	if len(v12) != len(v22) {
		return false
	}

	for _, v1 := range s1 {
		if v12[v1] != v22[v1] {
			return false
		}
	}

	return true
}

//!-
