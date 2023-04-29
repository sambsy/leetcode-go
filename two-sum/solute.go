package solute

// [两数之和](https://leetcode.cn/problems/two-sum/)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, value := range nums {
		r := target - value
		if v, ok := m[r]; ok {
			return []int{v, idx}
		}
		m[value] = idx
	}
	return nil
}
