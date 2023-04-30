package solute

// [移动石子直到连续](https://leetcode.cn/problems/moving-stones-until-consecutive/)
func numMovesStones(a int, b int, c int) []int {
	a, b, c = sort(a, b, c)

	min, max := 0, 0
	if c-a > 2 {
		if b-a <= 2 || c-b <= 2 {
			min = 1
		} else {
			min = 2
		}
	}

	max = c - a - 2
	return []int{min, max}

}
func sort(a, b, c int) (int, int, int) {
	if a > b {
		a, b = b, a
	}

	if b > c {
		b, c = c, b
	}

	if a > b {
		a, b = b, a
	}

	return a, b, c
}
