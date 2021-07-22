package utils

import (
	"strconv"
	"strings"
)

func Min(elements ...int) int {
	minimum := elements[0]
	for _, el := range elements {
		if minimum > el {
			minimum = el
		}
	}
	return minimum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(tree string) *TreeNode {
	v := strings.TrimSuffix(strings.TrimPrefix(tree, "["), "]")
	values := strings.Split(v, ",")
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 0, len(values))
	for _, val := range values {
		nodeVal, err := strconv.Atoi(val)
		if err == nil {
			nodes = append(nodes, &TreeNode{Val: nodeVal})
		} else {
			nodes = append(nodes, nil)
		}
	}
	nodeIdx := 0
	childIdx := 1
	root := nodes[0]
	for nodeIdx < len(values) {
		node := nodes[nodeIdx]
		if node != nil {
			if childIdx < len(values) {
				node.Left = nodes[childIdx]
				childIdx++
			}
			if childIdx < len(values) {
				node.Right = nodes[childIdx]
				childIdx++
			}
		}
		nodeIdx++
	}
	return root
}
