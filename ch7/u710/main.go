//rm u710; goimports -v -w main.go; gofmt -w main.go; go build; ./u710
package main

import (
	"fmt"
	"sort"
)

type word []byte

func (p word) Len() int {
	return len(p)
}

func isSym(i int, p word) bool {
	if i < len(p)/2 {
		return p[i] == p[len(p)-i-1]
	}
	if i > len(p)/2 {
		return p[i] == p[len(p)/2+(len(p)/2-i)]
	}
	return true
}

func (p word) Less(i, j int) bool {
	return !isSym(i, p) || !isSym(j, p)
}

func (p word) Swap(i, j int) {}

func IsPalindrom(s sort.Interface) bool {
	sort.Sort(s)
	return sort.IsSorted(s)
}

func main() {
	w := word("arozaupalanalapuazora")
	fmt.Printf("\nIs %s palindrom: %t\n", w, IsPalindrom(w))
}
