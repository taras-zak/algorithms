package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return helper(root, root.Val)
}

func helper(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	if root.Val >= max {
		max = root.Val
		return helper(root.Left, max) + helper(root.Right, max) + 1
	}
	return helper(root.Left, max) + helper(root.Right, max)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
			Left:  &TreeNode{Val: 2},
		},
	}
	fmt.Println(goodNodes(root))
}
