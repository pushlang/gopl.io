// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	i := 1
	for ; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	
	fmt.Printf("%d, %s", i, s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("\n%s", os.Args[1:])
}

//!-
