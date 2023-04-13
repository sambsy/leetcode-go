package solute

// [出现最频繁的偶数元素](https://leetcode.cn/problems/most-frequent-even-element/)
func mostFrequentEven(nums []int) int {
	ctrs := make(map[int]int)
	result := -1
	current := -1
	for _, num := range nums {
		if num%2 == 1 {
			continue
		}
		count, exist := ctrs[num]
		if !exist {
			ctrs[num] = 1
		} else {
			ctrs[num] += 1
		}
		count++

		if count > current {
			current = count
			result = num
		} else if count == current && num < result {
			result = num
		}
	}
	return result
}
