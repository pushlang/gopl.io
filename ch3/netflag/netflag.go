// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 77.

// Netflag demonstrates an integer type used as a bit field.
package main

import (
	"fmt"
	"net"
)

//!+
func isUp(v net.Flags) bool     { return v&net.FlagUp == net.FlagUp }
func turnDown(v *net.Flags)     { *v &^= net.FlagUp }
func setBroadcast(v *net.Flags) { *v |= net.FlagBroadcast }
func isCast(v net.Flags) bool   { return v&(net.FlagBroadcast|net.FlagMulticast) != 0 }

func main() {
	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, isUp(v)) // "10001 true"
	turnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v)) // "10000 false"
	setBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, isCast(v)) // "10010 true"
}

//!-
