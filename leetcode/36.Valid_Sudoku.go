package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	hash := make(map[string]struct{})
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			ch := string(board[row][col])
			if ch != "." {
				if _, ok := hash[ch+"row"+strconv.Itoa(row)]; ok {
					return false
				} else {
					hash[ch+"row"+strconv.Itoa(row)] = struct{}{}
				}
				if _, ok := hash[ch+"col"+strconv.Itoa(col)]; ok {
					return false
				} else {
					hash[ch+"col"+strconv.Itoa(col)] = struct{}{}
				}
				if _, ok := hash[ch+"box"+strconv.Itoa(row/3)+strconv.Itoa(col/3)]; ok {
					return false
				} else {
					hash[ch+"box"+strconv.Itoa(row/3)+strconv.Itoa(col/3)] = struct{}{}
				}
			}
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
	fmt.Println(isValidSudoku(board))
}
