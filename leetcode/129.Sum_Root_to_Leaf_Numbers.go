package main

import (
	"fmt"
	"github.com/taras-zak/algorithms/utils"
)

func sumNumbers(root *utils.TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *utils.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return sumNumbersHelper(root.Left, sum) + sumNumbersHelper(root.Right, sum)
}

func main() {
	fmt.Println(sumNumbers(utils.MakeTree("[1,2,3]")))
}
