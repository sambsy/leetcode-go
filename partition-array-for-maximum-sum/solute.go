package solute

// [分隔数组以得到最大和](https://leetcode.cn/problems/partition-array-for-maximum-sum/)
func maxSumAfterPartitioning(arr []int, k int) int {
	result := make([]int, 0, len(arr))
	for i := 0; i < k; i++ {
		result = append(result, maxSum(arr[:i+1]))
	}

	for j := k; j < len(arr); j++ {
		result = append(result, 0)
		for i := 0; i < k; i++ {
			mm := result[j-i-1] + maxSum(arr[j-i:j+1])
			if result[j] < mm {
				result[j] = mm
			}
		}
	}

	return result[len(result)-1]

}

func maxSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for j := 1; j < len(arr); j++ {
		if arr[j] > max {
			max = arr[j]
		}
	}

	return max * len(arr)
}
