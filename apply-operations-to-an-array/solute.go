package solute

// [对数组执行操作](https://leetcode.cn/problems/apply-operations-to-an-array/)
func applyOperations(nums []int) []int {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i+1 < len(nums) && nums[i] == nums[i+1] {
				nums[i-zeroCount] = nums[i] * 2
				zeroCount++
				i++
			} else {
				nums[i-zeroCount] = nums[i]
			}
		} else {
			zeroCount++
		}
	}

	for i := 0; i < zeroCount; i++ {
		nums[len(nums)-1-i] = 0
	}

	return nums
}
