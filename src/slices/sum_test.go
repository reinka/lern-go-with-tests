package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers)

		if want != got {
			t.Errorf("Got %d but want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	assertCorrectSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but want %v", got, want)
		}
	}

	t.Run("Sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 9})
		want := []int{5, 9}

		assertCorrectSum(t, got, want)
	})

	t.Run("Sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 9})
		want := []int{0, 9}

		assertCorrectSum(t, got, want)
	})
}
