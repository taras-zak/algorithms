package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	height := 0
	return helper(root, &height)
}

func helper(root *TreeNode, height *int) bool {
	if root == nil {
		*height = -1
		return true
	}
	left := 0
	right := 0
	l := helper(root.Left, &left)
	r := helper(root.Right, &right)
	diff := left - right
	if l && r && (diff > -2 && diff < 2) {
		if left < right {
			left = right
		}
		*height = left + 1
		return true
	}
	return false
}
