package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		height []int
		want   int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for _, cc := range cases {
		got := trap(cc.height)
		if got != cc.want {
			t.Errorf("failed!! height[%v],want[%d],got[%d]", cc.height, cc.want, got)
		}
	}
}
