package main

import "fmt"

func forward(x int) {
	fmt.Print(x, " ")
	if x > 0 {
		forward(x - 1)
	}
}

func backward(x int) {
	if x > 0 {
		backward(x - 1)
	}
	fmt.Print(x, " ")
}

func a1(n int) {
	fmt.Print(n, " ")
	b1(n - 1)
}

func b1(n int) {
	fmt.Print(n, " ")
	if n > 0 {
		b1(n - 1)
	}
}

func a2(n int) {
	//fmt.Print(n, " ")
	b2(n - 1)
	fmt.Print(n, " ")
}

func b2(n int) {
	// fmt.Print(n, " ")
	if n < 10 {
		a2(n + 2)
	}
	fmt.Print(n, " ")
}

func loopImitationI(i, n int) {
	fmt.Print(i, ":")
	if i < n {
		loopImitationJ(0, i, n)
		fmt.Println()
		loopImitationI(i+1, n)
	}
}

func loopImitationJ(i, j, n int) {
	fmt.Print(i*j, " ")
	if i < n {
		loopImitationJ(i+1, j, n)
	}
	fmt.Print(i*j, " ")
}

func binary(x int) {
	c := x % 2
	x = x / 2

	if x > 0 {
		binary(x)
	}

	fmt.Print(c)
}

func main() {
	forward(5)
	fmt.Println("")
	backward(5)
	fmt.Println("")
	a1(5)
	fmt.Println("")
	a2(1)
	fmt.Println("")
	loopImitationI(0, 5)
	fmt.Println("")
	fmt.Print("10:")
	binary(10)
}
