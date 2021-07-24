package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left == right {
		return 1<<left + countNodes(root.Right)
	} else {
		return 1<<right + countNodes(root.Left)
	}
}

func getDepth(root *TreeNode) int {
	var res int
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

func main() {

}
