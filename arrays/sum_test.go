package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %q want %q given %q", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Two slices", func(t *testing.T) {
		got := sumAll([]int{1, 1}, []int{2, 2})
		want := []int{2, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	cases := map[string]struct {
		got  []int
		want []int
	}{
		"three elements": {got: sumAllTails([]int{1, 1, 1}, []int{2, 2, 2}), want: []int{2, 4}},
		"empty slice":    {got: sumAllTails([]int{}), want: []int{0}},
	}

	for _, tc := range cases {
		t.Run("testcase", func(t *testing.T) {
			checkSums(t, tc.got, tc.want)
		})
	}
}
