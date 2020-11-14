package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Counter interface {
	Count(io.Reader) (int, error)
	Increment()
	Value() int
}

type byteCounter int
type wordCounter int
type lineCounter int

type CountWriter interface {
	io.Writer
	Counter
}

type countWriter struct {
	r io.Reader
	w io.Writer
	Counter
}

func (cw *countWriter) Write(p []byte) (int, error) {
	cw.w.Write(p)
	return cw.Counter.Count(cw.r)
}

// type counters []Counter
// type countWriters []*countWriter
// func (cw *countWriter) Write(p []byte) (int, error) {
// 	cw.w.Write(p)
// 	n, err := 0, error(nil)
//
// 	for _, c := range cw.counters {
// 		n, err = c.Counter.Count(c.r)
// 	}
//
// 	return n, err
// }

// func (c *counters) Count(r io.Reader) (int, error) {
// 	return count(c, r, bufio.ScanBytes)
// }
// func (c *counters) Increment() { *c++ }
// func (c *counters) Value() int { return int(*c) }

func (c *byteCounter) Count(r io.Reader) (int, error) {
	return count(c, r, bufio.ScanBytes)
}
func (c *byteCounter) Increment() { *c++ }
func (c *byteCounter) Value() int { return int(*c) }

func (c *wordCounter) Count(r io.Reader) (int, error) {
	return count(c, r, bufio.ScanWords)
}
func (c *wordCounter) Increment() { *c++ }
func (c *wordCounter) Value() int { return int(*c) }

func (c *lineCounter) Count(r io.Reader) (int, error) {
	return count(c, r, bufio.ScanLines)
}
func (c *lineCounter) Increment() { *c++ }
func (c *lineCounter) Value() int { return int(*c) }

func count(c Counter, r io.Reader, split bufio.SplitFunc) (int, error) {
	var done = make(chan struct{})

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(r)

	scanner.Split(split)

	go func() {
		for scanner.Scan() {
			c.Increment()
		}
		done <- struct{}{}
	}()

	<-done
	return c.Value(), scanner.Err()
}

func NewByteCounter() Counter {
	c := byteCounter(0)
	return &c
}
func NewWordCounter() Counter {
	c := wordCounter(0)
	return &c
}
func NewLineCounter() Counter {
	c := lineCounter(0)
	return &c
}

func NewCountWriter(r io.Reader, w io.Writer, c Counter) CountWriter {
	return &countWriter{r, w, c}
}

func main() {
	c := NewLineCounter()

	var text = `Hello my dear friends
Well, here I am on record at last
And it feels so wonderful to be here with you`

	//n, _ := c.Count(strings.NewReader(text), nil)

	cw := NewCountWriter(strings.NewReader(text), os.Stdout, c)

	fmt.Fprintf(cw, text)
	fmt.Println()
	fmt.Println(cw.Value())
}

// type counts struct {
// 	*byteCounter
// 	*wordCounter
// 	*lineCounter
// }

// func (c *counts) Count(r io.Reader, w io.Writer) (int, error) {
// 	var done = make(chan struct{})
//
// 	var scanner *bufio.Scanner
// 	scanner = bufio.NewScanner(r)
// 	scanner.Split(split)
//
// 	go func() {
// 		for scanner.Scan() {
// 			if w != nil {
// 				w.Write(scanner.Bytes())
// 			}
// 			*(c.byteCounter)++
// 		}
// 		done <- struct{}{}
// 	}()
//
// 	<-done
// 	return int(*(c.byteCounter)), scanner.Err()
// }
