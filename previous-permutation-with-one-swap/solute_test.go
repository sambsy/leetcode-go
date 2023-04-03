package solute

import (
	"testing"
)

func TestSolute(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		// {[]int{3, 2, 1}, []int{3, 1, 2}},
		// {[]int{1, 1, 5}, []int{1, 1, 5}},
		// {[]int{1, 9, 4, 6, 7}, []int{1, 7, 4, 6, 9}},
		{[]int{3, 1, 1, 3}, []int{1, 3, 1, 3}},
	}

	for _, c := range cases {
		got := prevPermOpt1(c.input)
		if len(got) != len(c.want) {
			t.Errorf("input %v, got %v, want %v", c.input, got, c.want)
		}
	}
}
