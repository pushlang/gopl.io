// The sha256 command computes the SHA256 hash (an array) of a string.

package main

import "testing"

func Test_shaComp(t *testing.T) {
	type args struct {
		s1 [32]byte
		s2 [32]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test0", args{[32]byte{0b00000001}, [32]byte{0b00000001}}, 0},
		{"test4", args{[32]byte{0b00000000}, [32]byte{0b10101001}}, 4},
		{"test12", args{[32]byte{0b00000000, 0b10101010}, [32]byte{0b10101001, 0b01010101}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shaComp(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("shaComp() = %v, want %v", got, tt.want)
			}
		})
	}
}
