package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	total := sum(root, nil)
	var res int
	sum(root, func(subtreeSum int) {
		prod := (total - subtreeSum) * subtreeSum
		if prod > res {
			res = prod
		}
	})
	return res % (1_000_000_000 + 7)
}

func sum(root *TreeNode, helper func(subtreeSum int)) int {
	if root == nil {
		return 0
	}
	sumBtreeSum := root.Val + sum(root.Left, helper) + sum(root.Right, helper)
	if helper != nil {
		helper(sumBtreeSum)
	}
	return sumBtreeSum
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
			Left: &TreeNode{
				Val: 4,
			},
		},
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(maxProduct(tree))
}
