package main

import "fmt"

func main() {
	var a [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = i * j
		}
	}

	//fmt.Printf("%v\n", a)

	b := make([][]int, 5)

	for i := 0; i < len(b); i++ {
		c := make([]int, 5)
		b[i] = c

		for j := 0; j < len(c); j++ {
			b[i][j] = i * j
		}
	}

	b = b[:2]

	for i := 0; i < len(b); i++ {
		b[i] = b[i][:2]
	}

	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Printf("\n")
	}

	// fmt.Printf("%b\n", &b[0][0])
	// fmt.Printf("%b\n", &b[0][1])
	// fmt.Printf("%b\n", &b[0][2])
}
