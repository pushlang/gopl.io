package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	if len(nums) > 0 {
		nums[1] = 0
	}
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			v++
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	a, b, c, d := 89, 89, 90, 95
	find(a, b, c, d)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	find(89, nums...)	
	fmt.Println(nums)
}
