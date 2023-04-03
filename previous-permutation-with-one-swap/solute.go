package solute

// [交换一次的先前排列](https://leetcode.cn/problems/previous-permutation-with-one-swap/)
func prevPermOpt1(arr []int) []int {
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			for j := len(arr) - 1; j > i; j-- {
				if j > i+1 && arr[j] == arr[j-1] {
					continue
				}
				if arr[j] < arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					return arr
				}
			}
		}
	}
	return arr
}
