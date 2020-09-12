package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Printf("%s\n", tempconv.CToK(tempconv.BoilingC))
	fmt.Printf("%s\n", tempconv.CToK(tempconv.AbsoluteZeroC))
}
