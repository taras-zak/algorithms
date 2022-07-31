package main

func findMinHeightTrees(n int, edges [][]int) []int {
	var res []int
	if n < 2 {
		for i := 0; i < n; i++ {
			res = append(res, i)
		}
		return res
	}
	graph := make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		graph[edge[0]][edge[1]] = struct{}{}
		graph[edge[1]][edge[0]] = struct{}{}
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(graph[i]) == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		n -= len(queue)
		newQueue := make([]int, 0)
		for _, node := range queue {
			child := getFirst(graph[node])
			delete(graph[child], node)
			if len(graph[child]) == 1 {
				newQueue = append(newQueue, child)
			}
		}
		queue = newQueue
	}
	for _, node := range queue {
		res = append(res, node)
	}
	return res
}

func getFirst(m map[int]struct{}) int {
	for k := range m {
		return k
	}
	return -1
}
