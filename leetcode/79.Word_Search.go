package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false

}

func dfs(grid [][]byte, i, j int, word string, curLetter int) bool {
	if curLetter == len(word) {
		return true
	}
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != word[curLetter] {
		return false
	}

	tmp := grid[i][j]
	grid[i][j] = '#'
	res := dfs(grid, i-1, j, word, curLetter+1) ||
		dfs(grid, i+1, j, word, curLetter+1) ||
		dfs(grid, i, j-1, word, curLetter+1) ||
		dfs(grid, i, j+1, word, curLetter+1)
	grid[i][j] = tmp
	return res
}
