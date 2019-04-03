package algorithms


func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix)==0 || len(matrix[0]) == 0 {
		return nil
	}
	i, j, k := 0, 0, 0
	isUp := true
	rows := len(matrix)
	cols := len(matrix[0])
	res := make([]int, cols*rows)
	for k < cols*rows {
		if isUp {
			for ; i >= 0 && j < cols; i,j = i-1, j+1 {
				res[k] = matrix[i][j]
				k += 1
			}
			if i == -1 && j <= cols-1 {
				i = 0
			}
			if j == cols {
				i = i + 2
				j -= 1
			}
		} else {
			for ; i < rows && j >= 0; i,j = i+1, j-1 {
				res[k] = matrix[i][j]
				k += 1
			}
			if j == -1 && i <= rows -1 {
				j = 0
			}
			if i == rows {
				j += 2
				i -= 1
			}
		}
		isUp = !isUp
	}
	return res
}

