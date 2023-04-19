package solute

import "testing"

func TestMaxSumAfterPartitioning(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{1, 15, 7, 9, 2, 5, 10}, 3, 84},
		{[]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4, 83},
		{[]int{1}, 1, 1},
	}
	for _, tt := range tests {
		if got := maxSumAfterPartitioning(tt.arr, tt.k); got != tt.want {
			t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
		}

	}
}
