// Because we are doing simple addition in these tests, a simple construct for iterating is key.
//If we were to add in latency within each call, the concurrent channel iterators would perform much better.
//Concurrency is a powerful thing, especially in the right context.
package iterators

import "testing"

func BenchmarkLoop(b *testing.B) {
	i := 10
	for n := 0; n < b.N; n++ {
		simpleforloop(i)
	}
}

func BenchmarkCallback(b *testing.B) {
	i := 10
	for n := 0; n < b.N; n++ {
		CallbackLoop(i)
	}
}

func BenchmarkNextLoop(b *testing.B) {
	i := 10
	for n := 0; n < b.N; n++ {
		NextLoop(i)
	}
}

func BenchmarkBufferedChan(b *testing.B) {
	i := 10
	for n := 0; n < b.N; n++ {
		BufferedChanLoop(i)
	}
}
