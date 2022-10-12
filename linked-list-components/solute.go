package solute

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// [题目信息](https://leetcode-cn.com/problems/linked-list-components/)
func numComponents(head *ListNode, nums []int) int {
	sort.Ints(nums)
	count := 0
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		if find(cur.Val, nums) {
			length++
		} else {
			if length > 0 {
				count++
				length = 0
			}
		}
	}
	if length > 0 {
		count++
	}

	return count
}

func find(val int, nums []int) bool {
	begin, end := 0, len(nums)
	for begin < end {
		if val > nums[end-1] || val < nums[begin] {
			return false
		}
		mid := (begin + end) / 2
		if midVal := nums[mid]; midVal == val {
			return true
		} else if midVal > val {
			end = mid
		} else {
			begin = mid
		}
	}

	return false
}
