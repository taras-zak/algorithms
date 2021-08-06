package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*Node{root}
	for len(queue) != 0 {
		levelCount := len(queue)
		level := make([]int, 0, levelCount)
		for levelCount != 0 {
			top := queue[0]
			queue = queue[1:]
			for _, child := range top.Children {
				queue = append(queue, child)
			}
			levelCount--
			level = append(level, top.Val)
		}
		res = append(res, level)
	}
	return res
}

func main() {
}
