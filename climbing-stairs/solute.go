package solute

// [爬楼梯](https://leetcode.cn/problems/climbing-stairs/)
func climbStairs(n int) int {
	left, right := 1, 1
	for i := 1; i < n; i++ {
		left, right = right, left+right
	}

	return right
}
