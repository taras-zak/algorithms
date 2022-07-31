package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	longestPath(root)
	return diameter
}

func longestPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := longestPath(root.Left)
	right := longestPath(root.Right)
	if diameter < left+right {
		diameter = left + right
	}
	if left < right {
		left = right
	}
	return left + 1
}
