package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

var Log = 14

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	depth := make(map[*TreeNode]int)
	up := make(map[*TreeNode][]*TreeNode)
	dfs(root, depth, up)
	if depth[p] < depth[q] {
		p, q = q, p
	}
	for depth[p] > depth[q] {
		p = up[p][0]
	}
	for p != q {
		p = up[p][0]
		q = up[q][0]
	}
	return p
}

func dfs(root *TreeNode, depth map[*TreeNode]int, up map[*TreeNode][]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		depth[root.Left] = depth[root] + 1
		up[root.Left] = []*TreeNode{root}
		for j := 1; j < Log; j++ {
			up[root.Left] = append(up[root.Left], up[root.Left][j-1])
		}
		dfs(root.Left, depth, up)
	}
	if root.Right != nil {
		depth[root.Right] = depth[root] + 1
		up[root.Right] = []*TreeNode{root}
		for j := 1; j < Log; j++ {
			up[root.Right] = append(up[root.Right], up[root.Right][j-1])
		}
		dfs(root.Right, depth, up)
	}
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(lowestCommonAncestor2(tree, tree.Left, tree.Right).Val)
}
