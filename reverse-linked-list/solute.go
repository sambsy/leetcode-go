package solute

type ListNode struct {
	Val  int
	Next *ListNode
}

// [反转链表](https://leetcode.cn/problems/reverse-linked-list/)
func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		value := head
		head = head.Next
		value.Next = result
		result = value
	}

	return result
}
