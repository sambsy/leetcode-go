package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, cc := range cases {
		got := searchInsert(cc.input, cc.target)
		if got != cc.want {
			t.Errorf("input %v, target %v, got %v, want %v", cc.input, cc.target, got, cc.want)
		}
	}
}
