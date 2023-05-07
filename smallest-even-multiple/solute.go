package solute

// [最小公倍数](https://leetcode.cn/problems/smallest-even-multiple/submissions/)
func smallestEvenMultiple(n int) int {
	value := 2
	for vv := n; vv%value != 0; value, vv = vv%value, value {
	}

	return 2 * n / value
}
