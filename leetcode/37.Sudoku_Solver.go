package main

import "fmt"

func solveSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				continue
			}
			for ch := byte('1'); ch <= '9'; ch++ {
				if isValid(board, row, col, ch) {
					board[row][col] = ch
					if solveSudoku(board) {
						return true
					} else {
						board[row][col] = '.'
					}
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, ch byte) bool {
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3
	for i := 0; i < 9; i++ {
		if board[row][i] == ch {
			return false
		}
		if board[i][col] == ch {
			return false
		}
		if board[blockRow+i/3][blockCol+i%3] == ch {
			return false
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for _, row := range board {
		for _, col := range row {
			fmt.Print(string(col), " ")
		}
		fmt.Println()
	}
}
