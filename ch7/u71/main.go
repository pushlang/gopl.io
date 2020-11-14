// 7.1, 7.2
//rm u71; goimports -v -w main.go ./count/count.go; gofmt -w main.go ./count/count.go; go build; ./u71
//del u71.exe & goimports -v -w main.go ./count/count.go & gofmt -w main.go ./count/count.go & go build & u71.exe

package main

import (
	"fmt"
	"os"
	"strings"

	"gopl.io/ch7/u71/count"
)

func main() {
	bc := count.NewByteCounter()
	wc := count.NewWordCounter()
	lc := count.NewLineCounter()

	var text = `Hello my dear friends
Well, here I am on record at last
And it feels so wonderful to be here with you on my first album
I'm so happy
Aha! Happy go lucky me
I just go my way
Living everyday
`

	cwb := count.NewCountWriter(strings.NewReader(text), os.Stdout, bc)
	cww := count.NewCountWriter(strings.NewReader(text), os.Stdout, wc)
	cwl := count.NewCountWriter(strings.NewReader(text), os.Stdout, lc)

	acw := count.NewAllCountWriters()

	acw.Add(cwb, cww, cwl)

	fmt.Fprintf(&acw, text)
	fmt.Println()
	fmt.Println(&acw)
	fmt.Printf("[%s %s %s]\n", bc, wc, lc)
}
