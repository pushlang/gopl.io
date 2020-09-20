package main

import "fmt"

func main() {
	var lines = make(map[string]map[string]bool)
	var fnames1 = make(map[string]bool)
	var fnames2 = make(map[string]bool)

	lines["a"] = fnames1
	lines["b"] = fnames2
	fnames1["fa"] = true
	fnames2["fb"] = true

	fmt.Printf("%t\n", lines["a"]["fa"])
	fmt.Printf("%t\n", lines["a"]["fb"])
	fmt.Printf("%t\n", lines["b"]["fa"])
	fmt.Printf("%t\n", lines["b"]["fb"])

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	defaultBool = customBool          //not allowed

	const a = 2.5 // not allowed as constant 2.5 truncated to integer
	var b = 5
	c := b + a
}
