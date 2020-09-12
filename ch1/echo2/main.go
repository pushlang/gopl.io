// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
//package main
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//s := []string(os.Args)

	s := []string{"a", "b", "c"}

	fmt.Println(Test1(s))
	fmt.Println(Test2(s))
}

func Test1(s []string) string {
	startTime := time.Now()
	ss, sep := "", ""
	for i := 0; i < 1000000; i++ {
		ss=""
		for _, arg := range s[1:] {
			//fmt.Printf("%d: %s\n", i, arg)
			ss += sep + arg
			sep = " "
		}
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	return ss
}

func Test2(s []string) string {
	startTime := time.Now()

	var ss string

	for i := 0; i < 1000000; i++ {
		ss = strings.Join(s[1:], " ")
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	return ss
}

//!-
