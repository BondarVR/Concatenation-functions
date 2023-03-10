package pkg

import "testing"

func BenchmarkConcat(b *testing.B) {
	strs := GenerateStrings(40)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(strs)
	}
}

func BenchmarkConcatOptimized(b *testing.B) {
	strs := GenerateStrings(40)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		СoncatOptimized(strs)
	}
}
