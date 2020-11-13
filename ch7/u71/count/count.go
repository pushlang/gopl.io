package count

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Counter interface {
	Increment()
	Value() []int
}

type CountWriter interface {
	io.Writer
	Counter
}

type countwriter struct {
	io.Writer
	Counter
}

func (c *countwriter) Write(p []byte) (int, error) {
	_, err := c.Writer.Write(p)
	switch c.Counter.(type) {
	case *byteCounter:
		_, err = counting(c, p, bufio.ScanBytes)
	case *wordCounter:
		_, err = counting(c, p, bufio.ScanWords)
	case *lineCounter:
		_, err = counting(c, p, bufio.ScanLines)
	case *allCounter:
		_, err = counting(c, p, bufio.ScanBytes)
		_, err = counting(c, p, bufio.ScanWords)
		_, err = counting(c, p, bufio.ScanLines)
	}
	return len(p), err
}

func NewCounterBytes() (Counter, *byteCounter) {
	c := new(byteCounter)
	return c, c
}
func NewCounterWords() (Counter, *wordCounter) {
	c := new(wordCounter)
	return c, c
}
func NewCounterLines() (Counter, *lineCounter) {
	c := new(lineCounter)
	return c, c
}
func NewCounterAll() (Counter, *allCounter) {
	a, b, c := byteCounter(0), wordCounter(0), lineCounter(0)
	all := &allCounter{&a, &b, &c}
	return all, all
}

func NewCountWriter(c Counter, w io.Writer) CountWriter {
	return &countwriter{Counter: c, Writer: w}
}

//func CountingWriter(w io.Writer) (io.Writer, *countwrite) {
//	sw := &countwrite{Writer: w}
//	return sw, &(sw.countwrite)
//}

type allCounter struct {
	byteCount *byteCounter
	wordCount *wordCounter
	lineCount *lineCounter
}

// func (c *allCounter) Write(p []byte) (int, error) {
// //_, err := c.Writer.Write(p)
// fmt.Println("allCounter", *c)
// _, err := c.byteCount.Write(p)
// _, err = c.wordCount.Write(p)
// _, err = c.lineCount.Write(p)

// return len(p), err
// }
func (c *allCounter) Increment() {
	// switch c.Counter.(type) {
	// case *byteCounter:
	// c.byteCount.Increment()
	// case *wordCounter:
	// c.wordCount.Increment()
	// case *lineCounter:
	// c.lineCount.Increment()
	// }
}
func (c *allCounter) Value() []int {
	return []int{int(*c.byteCount), int(*c.wordCount), int(*c.lineCount)}
}

// func (c allCounter) String() string {
// return fmt.Sprintf("[%v %v %v]", *(c.byteCount), *(c.wordCount), *(c.lineCount))
// }

type byteCounter int
type wordCounter int
type lineCounter int

// func (c *byteCounter) Write(p []byte) (int, error) {
// n, err := counting(c, p, bufio.ScanBytes)
// fmt.Println("byteCounter", *c)
// return n, err
// }
func (c *byteCounter) Increment() { *c++; fmt.Println(*c) }
func (c *byteCounter) Value() []int {
	return []int{int(*c)}
}

// func (c *wordCounter) Write(p []byte) (int, error) {
// n, err := counting(c, p, bufio.ScanWords)
// fmt.Println("wordCounter", *c)
// return n, err
// }
func (c *wordCounter) Increment() { *c++; fmt.Println(*c) }
func (c *wordCounter) Value() []int {
	return []int{int(*c)}
}

// func (c *lineCounter) Write(p []byte) (int, error) {
// n, err := counting(c, p, bufio.ScanLines)
// fmt.Println("lineCounter", *c)
// return n, err
// }
func (c *lineCounter) Increment() { *c++; fmt.Println(*c) }
func (c *lineCounter) Value() []int {
	return []int{int(*c)}
}

func counting(c Counter, p []byte, split func([]byte, bool) (int, []byte, error)) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	for scanner.Scan() {
		c.Increment()
	}

	return c.Value()[0], scanner.Err()
}
