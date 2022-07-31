package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cell := grid[i][j]
			if cell == '1' {
				dfs_islands(grid, i, j)
				count++
			}

		}
	}
	return count
}

func dfs_islands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 'v'
	dfs_islands(grid, i-1, j)
	dfs_islands(grid, i+1, j)
	dfs_islands(grid, i, j-1)
	dfs_islands(grid, i, j+1)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
	fmt.Println(grid)
}
