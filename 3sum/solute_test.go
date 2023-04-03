package solute

import "testing"

func TestSumThree(t *testing.T) {
	cases := []struct {
		input []int
		want  [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, c := range cases {
		got := threeSum(c.input)
		if len(got) != len(c.want) {
			t.Errorf("input %v, got %v, want %v", c.input, got, c.want)
		}
	}
}
