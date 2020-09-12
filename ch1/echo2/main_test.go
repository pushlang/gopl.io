// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import "testing"

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			"k", "l", "m", "n", "o", "p", "q", "q", "r", "s", "t",
			"u", "w", "w", "x", "y", "z"})
	}
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test2([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			"k", "l", "m", "n", "o", "p", "q", "q", "r", "s", "t",
			"u", "w", "w", "x", "y", "z"})
	}
}
