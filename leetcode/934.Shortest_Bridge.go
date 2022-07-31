package main

import (
	"fmt"
)

type coord struct {
	i, j int
}

func shortestBridge(grid [][]int) int {
	var queue []coord
outer:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = dfs(grid, i, j, queue)
				break outer
			}
		}
	}
	step := 0
	for len(queue) != 0 {
		size := len(queue)
		for n := 0; n < size; n++ {
			top := queue[n]
			for _, dir := range []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				i, j := top.i+dir.i, top.j+dir.j
				if inBoundaries(grid, i, j) && grid[i][j] != -1 {
					if grid[i][j] == 1 {
						return step
					}
					queue = append(queue, coord{i, j})
					grid[i][j] = -1
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return -1
}

func inBoundaries(grid [][]int, i, j int) bool {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
		return false
	}
	return true
}

func dfs(grid [][]int, i, j int, queue []coord) []coord {
	if !inBoundaries(grid, i, j) || grid[i][j] != 1 {
		return queue
	}
	queue = append(queue, coord{i, j})
	grid[i][j] = 2
	queue = dfs(grid, i+1, j, queue)
	queue = dfs(grid, i-1, j, queue)
	queue = dfs(grid, i, j+1, queue)
	queue = dfs(grid, i, j-1, queue)
	return queue
}

func main() {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println(shortestBridge(grid))
}
