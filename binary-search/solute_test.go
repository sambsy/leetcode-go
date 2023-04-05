package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 13, -1},
	}

	for _, cc := range cases {
		got := search(cc.input, cc.target)
		if got != cc.want {
			t.Errorf("input %v, target %v, got %v, want %v", cc.input, cc.target, got, cc.want)
		}

	}
}
