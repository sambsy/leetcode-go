package solute

// [题目信息](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/)
func maxChunksToSorted(arr []int) int {
	count := 0
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			count++
		}
	}

	return count
}
