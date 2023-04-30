package solute

import "sort"

// [三数之和](https://leetcode.cn/problems/3sum/)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for t := 0; t < len(nums); t++ {
		if nums[t] > 0 {
			break
		}
		if t > 0 && nums[t] == nums[t-1] {
			continue
		}

		l, r := t+1, len(nums)-1
		for l < r {
			sum := nums[t] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[t], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
