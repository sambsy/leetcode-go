package solute

// [两数之和](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/)
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		value := numbers[left] + numbers[right]
		if value == target {
			return []int{left + 1, right + 1}
		} else if value < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
