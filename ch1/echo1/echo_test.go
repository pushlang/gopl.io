package echo

import "testing"

//!-test

//!+bench
func BenchmarkEchoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo()
	}
}
