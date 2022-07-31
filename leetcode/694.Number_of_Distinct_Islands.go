package main

import (
	"strings"
)

func numDistinctIslands(grid [][]int) int {
	islands := make(map[string]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cell := grid[i][j]
			if cell == 1 {
				var sb strings.Builder
				dfs_islands(grid, i, j, &sb, '0')
				islands[sb.String()] = struct{}{}
			}
		}
	}
	return len(islands)
}

func dfs_islands(grid [][]int, i, j int, sb *strings.Builder, dir byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return
	}
	grid[i][j] = 0
	sb.WriteByte(dir)
	dfs_islands(grid, i-1, j, sb, 'U')
	dfs_islands(grid, i+1, j, sb, 'D')
	dfs_islands(grid, i, j-1, sb, 'L')
	dfs_islands(grid, i, j+1, sb, 'R')
	sb.WriteByte('0')
}
