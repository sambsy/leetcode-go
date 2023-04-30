package solute

// [接雨水](https://leetcode.cn/problems/trapping-rain-water/)
func trap(height []int) int {
	max := 0
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	water := 0
	left, right := 0, len(height)-1
	for i := 1; i <= max; i++ {
		for height[left] < i {
			left++
		}
		for height[right] < i {
			right--
		}

		for j := left + 1; j < right; j++ {
			if height[j] < i {
				water += 1
			}
		}
	}
	return water
}
