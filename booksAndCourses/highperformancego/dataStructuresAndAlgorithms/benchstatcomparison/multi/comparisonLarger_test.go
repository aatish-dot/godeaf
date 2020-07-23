package comparisonLarger

import (
	"testing"
)

func Benchmark_comparisonLarger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarJSON()
	}
}
