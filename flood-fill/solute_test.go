package solute

import (
	"testing"
)

func TestSolute(t *testing.T) {
	cases := []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			1,
			1,
			2,
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for _, c := range cases {
		got := floodFill(c.image, c.sr, c.sc, c.color)
		for i := 0; i < len(got); i++ {
			for j := 0; j < len(got[i]); j++ {
				if got[i][j] != c.want[i][j] {
					t.Errorf("floodFill(%v,%d,%d,%d) == %v,want %v", c.image, c.sr, c.sc, c.color, got, c.want)
				}
			}
		}
	}
}
