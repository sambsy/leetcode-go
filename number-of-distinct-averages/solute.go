package solute

import "sort"

// [不同的平均值数目](https://leetcode.cn/problems/number-of-distinct-averages/)
func distinctAverages(nums []int) int {
	sort.Ints(nums)
	avgs := make(map[int]struct{})

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		avg := float64(nums[i]+nums[j]) / 2
		avgs[int(avg*10)] = struct{}{}
	}

	return len(avgs)
}

/*
对原数组进行排序，然后用双指针，最左、最右元素为最小、最大元素，取平均值
使用 map 来对均值进行去重计数
float64 的值不可比较，由于已知最多有一位小数，可以将 平均数 * 10 后取整数部分作为key
*/
