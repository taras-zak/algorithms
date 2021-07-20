package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(f, s *TreeNode) bool {
	if f == nil && s == nil {
		return true
	}
	if f == nil || s == nil {
		return false
	}
	return f.Val == s.Val && checkSymmetric(f.Left, s.Right) && checkSymmetric(f.Right, s.Left)
}
