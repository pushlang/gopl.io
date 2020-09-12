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
	os.Args = []string{"", "-10,00", "-100,00", "-1000,00", "-10000,00", "-100000,00", "-1000000,00", "-10000000,00", "-100000000,00"}
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma2(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "." + s[n-3:]
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

//!-
