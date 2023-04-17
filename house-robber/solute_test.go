package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		args []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}

	for _, cc := range cases {
		got := rob(cc.args)
		if got != cc.want {
			t.Errorf("error!! args[%v],got[%v],want[%v]", cc.args, got, cc.want)
		}
	}
}
