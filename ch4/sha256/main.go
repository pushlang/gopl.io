// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("The count of different bits:%d\n", shaComp(c1, c2))
}

func shaComp(s1, s2 [32]byte) int {
	r := 0
	for i := range s1 {
		for j := byte(0); j < 8; j++ {
			if (((s1[i] >> j) & 1) ^ ((s2[i] >> j) & 1)) == 1 {
				r++
			}
		}

	}
	return r
}
