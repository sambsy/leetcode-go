package solute

import "testing"

func Test_findLengthOfShortestSubarray(t *testing.T) {
	cases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 2, 3, 10, 4, 2, 3, 5},
			want: 3,
		}, {
			arr:  []int{5, 4, 3, 2, 1},
			want: 4,
		}, {
			arr:  []int{1, 2, 3},
			want: 0,
		}, {
			arr:  []int{16, 10, 0, 3, 22, 1, 14, 7, 1, 12, 15},
			want: 8,
		}, {
			arr:  []int{48, 88, 16, 27, 46, 2, 45, 51, 62, 68, 57, 64, 93, 10, 48, 46, 17, 56, 66, 32, 58, 45, 16, 59, 65, 4, 15, 69, 35, 45, 65, 32, 78, 29, 51, 90, 64, 64, 52, 92, 57, 68, 43, 90, 78, 94, 64, 72},
			want: 45,
		}, {
			arr:  []int{1, 2, 3, 3, 10, 1, 3, 3, 5},
			want: 2,
		},
	}

	for _, cc := range cases {
		if result := findLengthOfShortestSubarray(cc.arr); result != cc.want {
			t.Errorf("[%v],failed,want[%v],result[%v]", cc.arr, cc.want, result)
		}
	}
}
