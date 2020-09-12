// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.

// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.

package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate([]int{0, 1, 2, 3, 4, 5}, 5)
	}
}

func BenchmarkRotate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate2([6]int{0, 1, 2, 3, 4, 5}, 5)
	}
}
