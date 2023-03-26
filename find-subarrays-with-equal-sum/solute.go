package solute

// 和相等的子数组
// [题目信息](https://leetcode.cn/problems/find-subarrays-with-equal-sum/)
func findSubarrays(nums []int) bool {
	sum := make(map[int]struct{}, len(nums))
	for i := 1; i < len(nums); i++ {
		value := nums[i] + nums[i-1]
		if _, exist := sum[value]; exist {
			return true
		}
		sum[value] = struct{}{}
	}
	return false
}
