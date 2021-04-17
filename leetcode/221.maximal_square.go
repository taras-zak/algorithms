package main

import (
	"fmt"
	"github.com/taras-zak/algorithms/utils"
)

func PrintIntMatrix(i [][]int) {
	for _, row := range i {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}

func maximalSquare(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows+1)
	for row := 0; row <= rows; row++ {
		dp[row] = make([]int, cols+1)
	}
	var res int
	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			if matrix[row-1][col-1] == 1 {
				dp[row][col] = utils.Min(dp[row-1][col], dp[row][col-1], dp[row-1][col-1]) + 1
				if dp[row][col] > res {
					res = dp[row][col]
				}
			} else {
				dp[row][col] = 0
			}
		}
	}
	return res * res
}

func main() {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	maximalSquare(matrix)
}
