package solute

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [节点与其祖先之间的最大差值](https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/)
func maxAncestorDiff(root *TreeNode) int {
	return diff(root, root.Val, root.Val)
}

func diff(root *TreeNode, max, min int) int {
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}

	result := max - min

	if root.Left != nil {
		if dd := diff(root.Left, max, min); dd > result {
			result = dd
		}
	}

	if root.Right != nil {
		if dd := diff(root.Right, max, min); dd > result {
			result = dd
		}
	}
	return result
}
