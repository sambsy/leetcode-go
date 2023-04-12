package solute

// [轮转数组](https://leetcode.cn/problems/rotate-array/)
func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}
	revert(nums)
	revert(nums[:k])
	revert(nums[k:])
}

func revert(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
