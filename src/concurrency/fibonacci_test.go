package concurrency

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.Run("Simple fibonacci algo", func(t *testing.T) {
		got := Fibonacci(MaxFibNumber)

		if !reflect.DeepEqual(got, fibSeq40) {
			t.Errorf("Got %v, want %v", got, fibSeq40)
		}
	})
}
