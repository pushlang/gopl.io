package count

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// AllCountWriters
type allCountWriters []*countWriter

func (cws *allCountWriters) Write(p []byte) (int, error) {
	(*cws)[0].w.Write(p)

	n, err := 0, error(nil)

	for _, cw := range *cws {
		n, err = cw.Count(cw.r)
	}

	return n, err
}
func (cws *allCountWriters) Add(cw ...*countWriter) { *cws = append(*cws, cw...) }

func (cws *allCountWriters) String() string {
	var b bytes.Buffer
	b.WriteByte('[')
	l := len(*cws)
	for i, cw := range *cws {
		b.WriteString(cw.String())

		if i < l-1 {
			b.WriteByte(' ')
		}
	}
	b.WriteByte(']')

	return b.String()
}

// CountWriter
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

// Counters
type Counter interface {
	Count(io.Reader) (int, error)
	Increment()
	Value() int
	String() string
}
type byteCounter struct {
	c int
	s bufio.SplitFunc
}
type wordCounter struct {
	c int
	s bufio.SplitFunc
}
type lineCounter struct {
	c int
	s bufio.SplitFunc
}

// Counters interface implementations
func (c *byteCounter) Count(r io.Reader) (int, error) {
	return count(c, r, c.s)
}
func (c *byteCounter) Increment()     { c.c++ }
func (c *byteCounter) Value() int     { return int(c.c) }
func (c *byteCounter) String() string { return fmt.Sprintf("bytes:%d", c.c) }

func (c *wordCounter) Count(r io.Reader) (int, error) {
	return count(c, r, c.s)
}
func (c *wordCounter) Increment()     { c.c++ }
func (c *wordCounter) Value() int     { return int(c.c) }
func (c *wordCounter) String() string { return fmt.Sprintf("words:%d", c.c) }

func (c *lineCounter) Count(r io.Reader) (int, error) {
	return count(c, r, c.s)
}
func (c *lineCounter) Increment()     { c.c++ }
func (c *lineCounter) Value() int     { return int(c.c) }
func (c *lineCounter) String() string { return fmt.Sprintf("lines:%d", c.c) }

func count(c Counter, r io.Reader, s bufio.SplitFunc) (int, error) {
	var done = make(chan struct{})

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(r)

	scanner.Split(s)

	go func() {
		for scanner.Scan() {
			c.Increment()
		}
		done <- struct{}{}
	}()

	<-done
	return c.Value(), scanner.Err()
}

// Counters, CounterWriter, allCountWriters constructors
func NewByteCounter() Counter {
	return &byteCounter{0, bufio.ScanBytes}
}
func NewWordCounter() Counter {
	return &wordCounter{0, bufio.ScanWords}
}
func NewLineCounter() Counter {
	return &lineCounter{0, bufio.ScanLines}
}

func NewCountWriter(r io.Reader, w io.Writer, c Counter) *countWriter {
	return &countWriter{r, w, c}
}

func NewAllCountWriters() allCountWriters {
	cw := make(allCountWriters, 0, 10)
	return cw
}
