package main

import (
	"fmt"
)

func main() {
	isEmpty := make([]int, 0) // =0
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isEmpty) == 0, isEmpty == nil, isEmpty)
	
	isEmpty = []int{} // =0
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isEmpty) == 0, isEmpty == nil, isEmpty)
		
	var isNil []int // nil
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isNil) == 0, isNil == nil, isNil)
	
	isNil = []int(nil) // nil
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isNil) == 0, isNil2 == nil, isNil)
	
	
	isEmptyMap := make(map[string]int) // =0
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isEmptyMap) == 0, isEmptyMap == nil, isEmptyMap)
	v,ok = isEmptyMap["key"] 
	fmt.Printf("?key==0: %t, ?v==0: %t, ?ok==false: %t\n", isEmptyMap["key"] == 0, v == 0, ok == false)
	
	isEmptyMap = map[string]int{} // =0
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isEmptyMap) == 0, isEmptyMap == nil, isEmptyMap)
	v,ok = isEmptyMap["key"] 
	fmt.Printf("?key==0: %t, ?v==0: %t, ?ok==false: %t\n", isEmptyMap["key"] == 0, v == 0, ok == false)
	
	var isNilMap map[string]int // =nil
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isNilMap) == 0, isNilMap == nil, isNilMap)
	v,ok = isNilMap["key"] 
	fmt.Printf("?key==0: %t, ?v==0: %t, ?ok==false: %t\n", isNilMap["key"] == 0, v == 0, ok == false)
	
	isNilMap = map[string]int(nil) // =nil
	fmt.Printf("?==0: %t, ?==nil: %t, type: %T\n", len(isNilMap) == 0, isNilMap == nil, isNilMap)
	v,ok = isNilMap["key"] 
	fmt.Printf("?key==0: %t, ?v==0: %t, ?ok==false: %t\n", isNilMap["key"] == 0, v == 0, ok == false)
}
