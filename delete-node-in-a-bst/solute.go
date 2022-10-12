package solute

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// deleteNode 删除节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return mergeTree(root.Left, root.Right)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func mergeTree(leftTree, rightTree *TreeNode) *TreeNode {
	if leftTree == nil {
		return rightTree
	}

	if rightTree == nil {
		return leftTree
	}

	newTree := mergeTree(leftTree, rightTree.Left)
	rightTree.Left = newTree
	return rightTree
}
