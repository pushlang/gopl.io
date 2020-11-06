// 7.1, 7.2
//rm u71; goimports -v -w *.go; gofmt -w *.go; go build; ./u71
//del u71.exe & goimports -v -w main.go & gofmt -w main.go & go build & u71.exe

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type byteCounter int
type wordCounter int
type lineCounter int

type allCounters struct {
	byteCount byteCounter
	wordCount wordCounter
	lineCount lineCounter
}

type ShellWriter struct {
	io.Writer
	allCounters
}

func (sw *ShellWriter) Write(p []byte) (int, error) {
	_, err := sw.Writer.Write(p)

	_, err = sw.byteCount.Write(p)
	_, err = sw.wordCount.Write(p)
	_, err = sw.lineCount.Write(p)

	return len(p), err
}

func CountingWriter(w io.Writer) (io.Writer, *allCounters) {
	sw := &ShellWriter{Writer: w}
	return sw, &(sw.allCounters)
}

type Counter interface {
	Increment()
	Value() int
}

func (c *byteCounter) Increment() { *c++ }
func (c *wordCounter) Increment() { *c++ }
func (c *lineCounter) Increment() { *c++ }

func (c *byteCounter) Value() int { return int(*c) }
func (c *wordCounter) Value() int { return int(*c) }
func (c *lineCounter) Value() int { return int(*c) }

func writeCount(c Counter, p []byte, split func([]byte, bool) (int, []byte, error)) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	for scanner.Scan() {
		c.Increment()
	}
	return c.Value(), scanner.Err()
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := writeCount(c, p, bufio.ScanBytes)
	return n, err
}

func (c *wordCounter) Write(p []byte) (int, error) {
	n, err := writeCount(c, p, bufio.ScanWords)
	return n, err
}

func (c *lineCounter) Write(p []byte) (int, error) {
	n, err := writeCount(c, p, bufio.ScanLines)
	return n, err
}

func main() {
	sw, c := CountingWriter(os.Stdout)

	var text = `Hello my dear friends
Well, here I am on record at last
And it feels so wonderful to be here with you`

	n, err := fmt.Fprintf(sw, "text: %s\n", text)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(*c, n)
}
