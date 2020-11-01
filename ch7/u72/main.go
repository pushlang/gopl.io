//rm u72; goimports -v -w *.go; gofmt -w *.go; go build; ./u72

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ShellWriter struct {
	baseWriter io.Writer
	countBytes int64
}

func (sw *ShellWriter) Write(p []byte) (int, error) {
	sw.baseWriter.Write(p)

	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		sw.countBytes++
	}
	return len(p), scanner.Err()
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var sw = new(ShellWriter)
	sw.baseWriter = w

	return sw, &(sw.countBytes)
}

func main() {
	sw, c := CountingWriter(os.Stdout)

	var text = `Hello my dear friends
Well, here I am on record at last
And it feels so wonderful to be here with you`

	fmt.Fprintf(sw, "text: %s\n", text)

	fmt.Println(*c)
}
