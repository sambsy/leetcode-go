package solute

type ListNode struct {
	Val  int
	Next *ListNode
}

// [删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	var (
		prev   *ListNode
		target = head
	)
	for p != nil {
		p = p.Next
		prev = target
		target = target.Next
	}

	if prev == nil {
		return target.Next
	}

	prev.Next = target.Next
	return head
}
