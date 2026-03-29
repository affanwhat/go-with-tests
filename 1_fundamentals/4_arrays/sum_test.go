package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	CheckSum := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Return sum of an array with 5 ints", func(t *testing.T) {
		nums := []int{10, 2, 5, 4, 3}
		sum := Sum(nums)
		expected := 24
		CheckSum(t, sum, expected)
	})

	t.Run("Return 0 for an array with 0 in its value", func(t *testing.T) {
		nums := []int{0}
		sum := Sum(nums)
		expected := 0
		CheckSum(t, sum, expected)
	})

	t.Run("Return 0 for an empty array", func(t *testing.T) {
		nums := []int{}
		sum := Sum(nums)
		expected := 0
		CheckSum(t, sum, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Return the sum of each array in the array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		CheckSums(t, got, want)
	})

	t.Run("Return 0 if the array is empty", func(t *testing.T) {
		got := SumAll(nil)
		want := []int{0}
		CheckSums(t, got, want)
	})

}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of the slices in an array", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 10})
		want := []int{5, 19}
		CheckSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		CheckSums(t, got, want)
	})
}

func CheckSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
