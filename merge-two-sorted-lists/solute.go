package solute

type ListNode struct {
	Val  int
	Next *ListNode
}

// [合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		result  *ListNode
		current *ListNode
	)

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if current == nil {
				current = list1
				result = current
			} else {
				current.Next = list1
				current = current.Next
			}
			list1 = list1.Next
		} else {
			if current == nil {
				current = list2
				result = current
			} else {
				current.Next = list2
				current = current.Next
			}
			list2 = list2.Next
		}
	}

	if list1 != nil {
		if current == nil {
			current = list1
			result = current
		} else {
			current.Next = list1
		}
	}
	if list2 != nil {
		if current == nil {
			current = list2
			result = current
		} else {
			current.Next = list2
		}
	}

	return result

}
