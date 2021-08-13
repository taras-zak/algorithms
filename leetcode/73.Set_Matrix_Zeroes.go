package main

import "fmt"

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	zeroFirstRow := false
	zeroFirstCol := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					zeroFirstRow = true
				}
				if j == 0 {
					zeroFirstCol = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if zeroFirstRow {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}
	if zeroFirstCol {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
	return
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
}
