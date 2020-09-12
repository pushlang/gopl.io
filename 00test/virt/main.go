package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("i=%d\n", getIndex115(a, 5))
	fmt.Printf("i=%d\n", getIndex116(a, 9))
	fmt.Printf("i=%d\n", getIndex115(a, 10))
	fmt.Printf("i=%d\n", getIndex116(a, 10))
	fmt.Printf("i=%d\n", getIndex117(a, 9))
	fmt.Printf("i=%d\n", getIndex117(a, 10))

}

func getIndex115(a []int, x int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		}
	}
	return 0
}

func getIndex116(a []int, x int) int {
	i := 0
	for a = append(a, x); a[i] != x; i++ {
	}
	if i > len(a)-2 {
		return 0
	}
	return i
}

func getIndex117(a []int, x int) int { //binary search a[k] != x && 
	i, j := 0, len(a)-1
	for k := (i + j) / 2; i <= j; k = (i + j) / 2 {
		switch {
		case x > a[k]:
			i = k + 1
		case x < a[k]:
			j = k - 1
		default:
			return k
		}
	}
	return 0
}
