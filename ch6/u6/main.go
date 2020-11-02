package main

import (
	"fmt"

	. "gopl.io/ch6/u6/intset"
)

func main() {
	is := new(IntSet)

	for i := 0; i < 4; i++ {
		is.Add(1 << i)
	}
	fmt.Printf("#1Add: %v, Len:%d\n", is, is.Len())

	for i := 0; i <= 8; i++ {
		fmt.Printf("%d=%t:", i, is.Has(i))
	}
	fmt.Println()

	is2 := new(IntSet)

	for i := 3; i < 8; i++ {
		is2.Add(2 << i)
	}
	fmt.Printf("#2Add: %v, Len:%d\n", is2, is2.Len())

	is.UnionWith(is2)
	fmt.Printf("UnionWith, is/is2:%v, Len:%d\n", is, is.Len())

	for i := 0; i < 4; i++ {
		fmt.Printf("Get:%d, Value:%d\n", i, is.GetValue(i))
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("Get:%d, Position:%d\n", 1<<i, is.GetPosition(1<<i))
	}

	is.Remove(8)
	fmt.Printf("Remove:%v, is:%v\n", 8, is)
	is.Remove(16)
	fmt.Printf("Remove:%v, is:%v\n", 16, is)

	is3 := is.Copy()

	is.Clear()
	fmt.Printf("Clear all, is:%v\n", is)
	fmt.Printf("Copy, is3:%v\n", is3)

	is.AddAll(1<<0, 1<<1, 1<<2, 1<<3, 1<<4)
	fmt.Printf("AddAll, is:%v\n", is)

	is.IntersectionWith(is3)
	fmt.Printf("IntersectionWith, is/is3:%v\n", is)

	is3.DifferenceWith(is)
	fmt.Printf("DifferenceWith, is3/is:%v\n", is3)

	is3.Add(1 << 9)
	is3.SymmDifferenceWith(is2)
	fmt.Printf("SymmDifferenceWith, is3/is:%v\n", is3)

	fmt.Printf("Elems, is2: ")
	for _, e := range is2.Elems() {
		fmt.Printf("%d, ", e)
	}
	fmt.Println()
}
