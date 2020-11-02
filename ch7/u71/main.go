//rm u71; goimports -v -w *.go; gofmt -w *.go; go build; ./u71

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

type AllCounters struct {
	byteCounter ByteCounter
	wordCounter WordCounter
	lineCounter LineCounter
}

type Counter interface {
	Increment()
	Value() int
}

func (c *ByteCounter) Increment() { *c++ }
func (c *WordCounter) Increment() { *c++ }
func (c *LineCounter) Increment() { *c++ }

func (c *ByteCounter) Value() int { return int(*c) }
func (c *WordCounter) Value() int { return int(*c) }
func (c *LineCounter) Value() int { return int(*c) }

func (c *AllCounters) Write(p []byte) (int, error) {
	_, err := c.byteCounter.Write(p)
	_, err = c.wordCounter.Write(p)
	_, err = c.lineCounter.Write(p)
	return len(p), err
}

func writeCounter(c Counter, p []byte, split func([]byte, bool) (int, []byte, error)) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	for scanner.Scan() {
		c.Increment()
	}
	return c.Value(), scanner.Err()
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := writeCounter(c, p, bufio.ScanBytes)
	return n, err
}

func (c *WordCounter) Write(p []byte) (int, error) {
	n, err := writeCounter(c, p, bufio.ScanWords)
	return n, err
}

func (c *LineCounter) Write(p []byte) (int, error) {
	n, err := writeCounter(c, p, bufio.ScanLines)
	return n, err
}

func main() {
	var c AllCounters

	var text = `Hello my dear friends
Well, here I am on record at last
And it feels so wonderful to be here with you`

	fmt.Fprintf(&c, "text: %s", text)
	fmt.Println(c)
}
