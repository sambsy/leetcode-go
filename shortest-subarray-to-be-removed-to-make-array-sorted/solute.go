package solute

import "sort"

// 删除最短的子数组使剩余数组有序
// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/g
func findLengthOfShortestSubarray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < len(arr)-1 && arr[l] <= arr[l+1] {
		l++
	}

	for r > 0 && arr[r] >= arr[r-1] {
		r--
	}

	if l >= r {
		return 0
	}

	if arr[l] <= arr[r] {
		return r - l - 1
	}
	result := min(len(arr)-l-1, r)
	rr := r
	for i := l; i >= 0; i-- {
		//for rr >= r && arr[rr] >= arr[i] {
		//	rr--
		//}
		rr := r + sort.SearchInts(arr[rr:], arr[i])
		result = min(result, rr-i-1)
		//rr = rr - 1
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
