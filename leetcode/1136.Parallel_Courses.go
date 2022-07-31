package main

import (
	"fmt"
)

func minimumSemesters(n int, relations [][]int) int {
	adjList := make(map[int][]int, n)
	indegree := make(map[int]int)
	for _, edge := range relations {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	var res []int
	var queue []int
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	step := 0
	for len(queue) != 0 {
		step++
		var level []int
		for _, node := range queue {
			for _, k := range adjList[node] {
				indegree[k]--
				if indegree[k] == 0 {
					level = append(level, k)
				}
			}
			res = append(res, node)
		}
		queue = level
	}
	if len(res) != n {
		return -1
	}
	return step
}

func main() {
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
}
