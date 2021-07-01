package main

import (
	"fmt"
	"strconv"
	"strings"
)

const SPACE = 10

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (n *TreeNode) print(depth int) string {
	var r = ""
	if n.Right != nil {
		r += n.Right.print(depth + SPACE)
	}

	r += fmt.Sprintf("\n%s %d", strings.Repeat(" ", depth), n.Val)
	if n.Left != nil {
		r += n.Left.print(depth + SPACE)
	}
	return r
}

func (n *TreeNode) String() string {
	return n.print(0)
}

func (n *TreeNode) preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	res = append(res, n.preorderTraversal(root.Left)...)
	res = append(res, n.preorderTraversal(root.Right)...)
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			continue
		}
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func levelOrderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		res = append(res, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return res
}

func levelOrderNestedList(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		lvlCount := len(queue)
		r := make([]int, 0)
		for lvlCount > 0 {
			top := queue[0]
			queue = queue[1:]
			if top != nil {
				r = append(r, top.Val)
				if top.Left != nil {
					queue = append(queue, top.Left)
				}
				if top.Right != nil {
					queue = append(queue, top.Right)
				}
			}
			lvlCount--
		}
		res = append(res, r)
	}
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

func hasPathSum(root *TreeNode, sum int) bool {
	switch {
	case root == nil:
		return false
	case root.Left == nil && root.Right == nil && root.Val == sum:
		return true
	case hasPathSum(root.Left, sum-root.Val):
		return true
	case hasPathSum(root.Right, sum-root.Val):
		return true
	default:
		return false
	}
}

func printTreeVertical(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	depth := maxDepth(root)
	s := 1<<uint(depth) - 1
	res := make([][]string, 0)
	for i := 0; i < depth; i++ {
		//res = append(res, make([]string, s))
		row := make([]string, s)
		for r := range row {
			row[r] = " "
		}
		res = append(res, row)
	}
	fill(root, res, 0, 0, s+1)
	return res
}

func fill(root *TreeNode, m [][]string, i, l, r int) {
	if root == nil {
		return
	}
	m[i][(l+r-1)/2] = strconv.Itoa(root.Val)
	fill(root.Left, m, i+1, l, (l+r)/2)
	fill(root.Right, m, i+1, (l+r+1)/2, r)

}
