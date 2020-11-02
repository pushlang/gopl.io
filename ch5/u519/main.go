package main

import (
	"fmt"
)

func retValue(i int) {
	panic(i+1)
}

func main() {
	defer func() {
		if p:= recover(); p!= nil {
			fmt.Println("returned value:", p)
		}
	}()
	
	retValue(1)
}
