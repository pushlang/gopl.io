//del u.exe & goimports -v -w main.go & gofmt -w main.go & go build & u.exe

package main

import (
	"fmt"

	t "gopl.io/ch7/u/test"
)

func main() {
	a := t.New()
	fmt.Printf("%T\n", a)
	fmt.Printf("%v %v", a.F(), a.G())
}
