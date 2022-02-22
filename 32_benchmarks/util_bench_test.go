package util

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(35)
	}
}

func BenchmarkMakeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSlice(3000000)
	}
}
