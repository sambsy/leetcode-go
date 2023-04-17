package solute

// [打家劫舍](https://leetcode.cn/problems/house-robber/)
func rob(nums []int) int {
	storage := make([]int, 0, len(nums))
	storage = append(storage, nums[0])

	if len(nums) > 1 {
		storage = append(storage, max(nums[0], nums[1]))
	}

	for j := 2; j < len(nums); j++ {
		storage = append(storage, max(storage[j-1], storage[j-2]+nums[j]))
	}

	return storage[len(storage)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
