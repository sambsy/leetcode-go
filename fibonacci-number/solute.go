package solute

// [斐波那契数](https://leetcode.cn/problems/fibonacci-number/)
func fib(n int) int {
	if n == 0 {
		return 0
	}
	left, right := 0, 1
	for i := 2; i <= n; i++ {
		left, right = right, left+right
	}

	return right
}
