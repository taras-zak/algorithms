package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TODO: test traversals
// TODO: Bench iter vs recur: mem cpu

func TestBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}
	//fmt.Println(tree)

	tree = &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				&TreeNode{7, nil, nil},
				nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6,
				&TreeNode{8, nil, nil},
				&TreeNode{9, nil, nil}},
			nil}}
	PrettyPrint(printTreeVertical(tree))
}

func PrettyPrint(v [][]string) {
	b := bytes.NewBufferString("")
	for _, row := range v {
		b.WriteString(strings.Join(row, ""))
		b.WriteString("\n")
	}
	fmt.Println(b.String())
	return
}
