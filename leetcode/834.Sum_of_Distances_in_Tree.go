package main

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ans := make([]int, n)
	count := make([]int, n)
	tree := make([][]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}
	isSeen := make([]bool, n)
	postOrder(0, tree, count, ans, isSeen)
	isSeen = make([]bool, n)
	preOrder(0, tree, count, ans, isSeen)
	return ans
}

func postOrder(root int, tree [][]int, count, ans []int, isSeen []bool) {
	isSeen[root] = true
	for _, child := range tree[root] {
		if isSeen[child] {
			continue
		}
		postOrder(child, tree, count, ans, isSeen)
		count[root] += count[child]
		ans[root] += ans[child] + count[child]
	}
	count[root]++
}

func preOrder(root int, tree [][]int, count, ans []int, isSeen []bool) {
	isSeen[root] = true
	for _, child := range tree[root] {
		if isSeen[child] {
			continue
		}
		ans[child] = ans[root] - count[child] + len(count) - count[child]
		preOrder(child, tree, count, ans, isSeen)
	}
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}))
}
