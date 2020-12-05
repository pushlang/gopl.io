package main

import (
	"fmt"
)

func s(a struct{ a, b int }) {
	fmt.Printf("%T: %v\n", a, a)
}

func main() {
	type AB1 struct{ a, b int }
	type AB2 struct{ a, b int }
	//type CD struct{ c, d int }

	ab1 := AB1{1, 2}
	ab2 := AB2{1, 2}
	//cd := CD{1, 2}

	abl := struct{ a, b int }{1, 2}

	fmt.Println(abl == ab1, abl == ab2, ab1 == AB1(ab2))

	//fmt.Println(ab1 == AB1(cd)) // mismatched types

	s(ab1)
	s(ab2)
}
