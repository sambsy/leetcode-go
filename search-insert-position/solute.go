package solute

// [搜索插入位置](https://leetcode.cn/problems/search-insert-position/)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if right < 0 {
		return 0
	}
	if left >= len(nums) {
		return len(nums)
	}

	return left
}
