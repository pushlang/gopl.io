//rm u710; goimports -v -w main.go; gofmt -w main.go; go build; ./u710
package main

import (
	"fmt"
	"sort"
)

type word []byte

func (p word) Len() int {
	return 2
}

func (p word) Less(i, j int) bool {
	hl := len(p) / 2
	for i, v := range p {
		if i != hl {
			if v != p[2*hl-i] {
				return true
			}
		}
	}
	return false
}

func (p word) Swap(i, j int) {}

func IsPalindrom(s sort.Interface) bool {
	sort.Sort(s)
	return sort.IsSorted(s)
}

func main() {
	w := word("arozaupalanalapuazora")
	w2 := word("abb")
	w3 := word("aaa")

	fmt.Printf("\nIs %s palindrom: %t\n", w, IsPalindrom(w))
	fmt.Printf("\nIs %s palindrom: %t\n", w2, IsPalindrom(w2))
	fmt.Printf("\nIs %s palindrom: %t\n", w3, IsPalindrom(w3))
}
