package main

func findDiagonalOrder(nums [][]int) []int {
	diagonals := make(map[int][]int)
	for i, row := range nums {
		for j, col := range row {
			diagonals[i+j] = append(diagonals[i+j], col)
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(diagonals); i++ {
		if i%2 == 1 {
			res = append(res, diagonals[i]...)
		} else {
			for j := len(diagonals[i]) - 1; j >= 0; j-- {
				res = append(res, diagonals[i][j])
			}
		}
	}
	return res
}
