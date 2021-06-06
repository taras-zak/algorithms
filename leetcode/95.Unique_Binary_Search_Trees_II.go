package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return genTrees(1, n)
}

func genTrees(s, e int) []*TreeNode {
	if s > e {
		return []*TreeNode{nil}
	}
	if s == e {
		return []*TreeNode{{Val: s}}
	}
	res := make([]*TreeNode, 0)
	for i := s; i <= e; i++ {
		for _, lNode := range genTrees(s, i-1) {
			for _, rNode := range genTrees(i+1, e) {
				res = append(res, &TreeNode{Val: i, Left: lNode, Right: rNode})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(len(generateTrees(4)))
}
