package chapter4

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("adding collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTrails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTrailsNew(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 4}, []int{5, 6})
		want := []int{4, 6}

		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3})
		want := []int{0, 3}

		checkSums(t, got, want)
	})
}
