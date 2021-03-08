package concurrency

import (
	"reflect"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	b.Run("Sequential Fibonacci", func(b *testing.B) {
		got := Fibonacci(MaxFibNumber)

		if !reflect.DeepEqual(got, fibSeq40) {
			b.Errorf("Got %v, want %v", got, fibSeq40)
		}
	})

	b.Run("Concurrent Fibonacci", func(b *testing.B) {
		got := FibonacciConcurrent(MaxFibNumber)

		if !reflect.DeepEqual(got, fibSeq40) {
			b.Errorf("got %v, want %v", got, fibSeq40)
		}
	})
}
