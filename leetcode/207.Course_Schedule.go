package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	incEdges := make([]int, numCourses)
	for _, edge := range prerequisites {
		incEdges[edge[0]]++
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if incEdges[i] == 0 {
			queue = append(queue, i)
		}
	}
	edgeCnt := len(prerequisites)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for _, edg := range adjList[top] {
			edgeCnt--
			incEdges[edg]--
			if incEdges[edg] == 0 {
				queue = append(queue, edg)
			}
		}
	}
	return edgeCnt == 0
}
