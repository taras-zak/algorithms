package main

// DFS topological sort
func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	for _, edge := range prerequisites {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
	}
	visited := make(map[int]struct{})
	tempMark := make(map[int]struct{})
	labels := make([]int, 0)
	for node := numCourses - 1; node >= 0; node-- {
		if _, ok := visited[node]; !ok {
			labels, ok = DFSTopo(adjList, node, visited, tempMark, labels)
			if !ok {
				return nil
			}
		}
	}
	return labels
}

func DFSTopo(g map[int][]int, start int, visited, tempMark map[int]struct{}, labels []int) ([]int, bool) {
	visited[start] = struct{}{}
	tempMark[start] = struct{}{}
	for _, neighbour := range g[start] {
		if _, ok := visited[neighbour]; !ok {
			labels, ok = DFSTopo(g, neighbour, visited, tempMark, labels)
			if !ok {
				return nil, ok
			}
		} else if _, ok := tempMark[neighbour]; ok {
			return nil, false
		}
	}
	delete(tempMark, start)
	labels = append(labels, start)
	return labels, true
}
