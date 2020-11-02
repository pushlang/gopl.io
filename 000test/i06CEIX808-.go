package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr2d := make([][]string, 10)
	for i, v2d := range arr2d {
		v2d = make([]string, i+1)
		for j, v := range v2d {
			v = strconv.Itoa(i * j)

			fmt.Printf("%d:%d = %s, ", i, j, v)
		}

		fmt.Println()
	}
}
