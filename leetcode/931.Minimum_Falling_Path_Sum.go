package main

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for row := 1; row < n; row++ {
		for col := 0; col < n; col++ {
			min := matrix[row-1][col]
			if col-1 >= 0 && matrix[row-1][col-1] < min {
				min = matrix[row-1][col-1]
			}
			if col+1 < n && matrix[row-1][col+1] < min {
				min = matrix[row-1][col+1]
			}
			matrix[row][col] = min + matrix[row][col]
		}
	}
	min := n * 100
	for col := 0; col < n; col++ {
		if matrix[n-1][col] < min {
			min = matrix[n-1][col]
		}
	}
	return min
}
