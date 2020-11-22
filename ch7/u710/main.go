// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type word []byte

func (p word) Len() int           { return len(p) }
func (p word) Less(i, j int) bool { 
	return !s.Less(i ,j) && !s.Less(j, i) 
}
func (p word) Swap(i, j int) {}

func IsPalindrom(s sort.Interface) bool {
	sort.Sort(s)
	return sort.IsSorted(s)
}

func main() {
	w := word{"arosaupalanalapuazora"}
	IsPalindrom(word)
}
