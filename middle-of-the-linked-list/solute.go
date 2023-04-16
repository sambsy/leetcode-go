package solute

type ListNode struct {
	Val  int
	Next *ListNode
}

// [链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/)
func middleNode(head *ListNode) *ListNode {
	count := 0

	for p := head; p != nil; p = p.Next {
		count++
	}

	mid := count / 2
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}
