package solute

// [移动零](https://leetcode.cn/problems/move-zeroes/)
func moveZeroes(nums []int) {
	for left, right := -1, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			if left == -1 {
				left = right
			}
		} else if left >= 0 {
			nums[left], nums[right] = nums[right], 0
			left++
		}
	}
}

func moveZeroes2(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}

	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}
