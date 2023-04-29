package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}, 4},
	}
	for _, cc := range cases {
		got := longestConsecutive(cc.nums)
		if got != cc.want {
			t.Errorf("got %d want %d", got, cc.want)
		}
	}
}
