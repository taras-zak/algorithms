package main

import "fmt"

type pair struct {
	x, y int
}

func largestIsland(grid [][]int) int {
	index := 2
	res := 0
	areas := make(map[int]int)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				areas[index] = bfs(grid, pair{x, y}, index)
				res = max(res, areas[index])
				index++
			}
		}
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				size := 1
				neighbourAreas := make(map[int]struct{})
				for _, neighbourArea := range getNeighbours(pair{x, y}, len(grid)) {
					areaIdx := grid[neighbourArea.x][neighbourArea.y]
					if _, ok := neighbourAreas[areaIdx]; areaIdx > 1 && !ok {
						size += areas[areaIdx]
						neighbourAreas[areaIdx] = struct{}{}
					}
				}
				res = max(res, size)
			}
		}
	}
	return res
}

func bfs(grid [][]int, pos pair, index int) int {
	area := 0
	queue := []pair{pos}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if grid[top.x][top.y] == 1 {
			grid[top.x][top.y] = index
			area++
			for _, neighbour := range getNeighbours(top, len(grid)) {
				queue = append(queue, neighbour)
			}
		}
	}
	return area
}

func valid(pos pair, n int) bool {
	return 0 <= pos.x && pos.x < n && 0 <= pos.y && pos.y < n
}

func getNeighbours(pos pair, n int) []pair {
	res := make([]pair, 0)
	var neighbour pair
	for _, dir := range []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		neighbour = pair{dir.x + pos.x, dir.y + pos.y}
		if valid(neighbour, n) {
			res = append(res, neighbour)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestIsland([][]int{
		{1, 0},
		{0, 1},
	}))
}
