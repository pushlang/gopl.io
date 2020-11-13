package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Counter interface {
	Count(io.ReadWriter, func([]byte, bool) (int, []byte, error)) (int, error)
}

type bytecounter int

var done = make(chan struct{})

func (c *bytecounter) Count(r io.Reader, split func([]byte, bool) (int, []byte, error)) (int, error) {
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(r)
	scanner.Split(split)

	countreads := func() {
		for scanner.Scan() {
			*c++
		}
		done <- struct{}{}
	}
	switch r.(type) {
	case io.Reader:
		go countreads()
	case io.Writer:
		//countwrites()
	}
	<-done
	return int(*c), scanner.Err()
}

func main() {
	a := new(bytecounter)

	n, _ := a.Count(strings.NewReader("abcdefgh"), bufio.ScanBytes)

	fmt.Println(n)
}
