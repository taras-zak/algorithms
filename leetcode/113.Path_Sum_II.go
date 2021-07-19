package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return append(res, []int{root.Val})
		}
		return res
	}
	for _, path := range pathSum(root.Left, sum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	for _, path := range pathSum(root.Right, sum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	return res
}

func main() {

}
