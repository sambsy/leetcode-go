package solute

// [公因数的数目](https://leetcode.cn/problems/number-of-common-factors/)
func commonFactors(a int, b int) int {
	num := 0
	less := a
	if a > b {
		less = b
	}

	for i := 1; i <= less; i++ {
		if a%i == 0 && b%i == 0 {
			num++
		}
	}

	return num
}
