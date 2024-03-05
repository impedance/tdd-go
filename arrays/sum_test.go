package main

import (
	"reflect"
	"testing"
)


func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given: %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {

  checkSums := func(t testing.TB, got, want []int) {
  // func checkSums(t testing.TB, got, want []int) {
    t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
  }
	t.Run("two las numbers", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

    checkSums(t, got, want)
	})

  t.Run("three numbers", func(t *testing.T) {


		got := SumAllTails([]int{1, 2, 4, 5}, []int{0, 9, 1})
		want := []int{11, 10}

    checkSums(t, got, want)
  })

  t.Run("sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}

    checkSums(t, got, want)
  })
}
