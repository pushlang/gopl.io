package main

import (
	"testing"
)

//!-test

//!+bench
func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma("10000000000000000000000000000000000000")
	}
}

func BenchmarkComma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma2("10000000000000000000000000000000000000")
	}
}

func BenchmarkCommaRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commaRec("10000000000000000000000000000000000000")
	}
}

func BenchmarkAna(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ana("abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkAna2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ana2("abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba")
	}
}

func Test_ana(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true0", args{"", ""}, true},
		{"true1", args{"abc", "abc"}, true},
		{"true2", args{"abc", "cba"}, true},
		{"true3", args{"aabbcc", "ccbbaa"}, true},
		{"false1", args{"abc", "cbf"}, false},
		{"false2", args{"abc", "cbaa"}, false},
		{"false3", args{"abbcc", "ccbbaa"}, false},
		{"false4", args{"aabbcc", ""}, false},
		{"false5", args{"", "aabbcc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ana(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("ana(%v, %v) = %v, want %v", tt.args.s1, tt.args.s2, got, tt.want)
			}
		})
	}
}

func Test_ana2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true0", args{"", ""}, true},
		{"true1", args{"abc", "abc"}, true},
		{"true2", args{"abc", "cba"}, true},
		{"true3", args{"aabbcc", "ccbbaa"}, true},
		{"false1", args{"abc", "cbf"}, false},
		{"false2", args{"abc", "cbaa"}, false},
		{"false3", args{"abbcc", "ccbbaa"}, false},
		{"false4", args{"aabbcc", ""}, false},
		{"false5", args{"", "aabbcc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ana2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("ana(%v, %v) = %v, want %v", tt.args.s1, tt.args.s2, got, tt.want)
			}
		})
	}
}
