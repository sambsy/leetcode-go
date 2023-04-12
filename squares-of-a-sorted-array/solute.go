package solute

// [有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/`)
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	result := make([]int, len(nums))
	idx := len(result) - 1
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[right] >= nums[left] {
			result[idx] = nums[right]
			right--
		} else {
			result[idx] = nums[left]
			left++
		}
		idx--
	}
	return result
}
