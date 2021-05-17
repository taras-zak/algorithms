package main

import "fmt"

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	MOD := 1_000_000_007
	dp := make([][][]int, maxMove+1)
	for i := 0; i <= maxMove; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for move := 1; move <= maxMove; move++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for _, dir := range dirs {
					di := i + dir[0]
					dj := j + dir[1]
					if di < 0 || di >= m || dj < 0 || dj >= n {
						dp[move][i][j] = (dp[move][i][j] + 1) % MOD
					} else {
						dp[move][i][j] = (dp[move][i][j] + dp[move-1][di][dj]) % MOD
					}
				}
			}
		}
	}
	return dp[maxMove][startRow][startColumn]
}

func findPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	var cache [50][50][51]*int
	return dfsFindPaths2(m, n, maxMove, startRow, startColumn, &cache)
}

func dfsFindPaths2(m, n, maxMove, startRow, startColumn int, cache *[50][50][51]*int) int {
	if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n {
		return 1
	}
	if maxMove == 0 {
		return 0
	}
	if cache[startRow][startColumn][maxMove] != nil {
		return *cache[startRow][startColumn][maxMove]
	}
	res := (dfsFindPaths2(m, n, maxMove-1, startRow-1, startColumn, cache)%1_000_000_007 +
		dfsFindPaths2(m, n, maxMove-1, startRow+1, startColumn, cache)%1_000_000_007 +
		dfsFindPaths2(m, n, maxMove-1, startRow, startColumn-1, cache)%1_000_000_007 +
		dfsFindPaths2(m, n, maxMove-1, startRow, startColumn+1, cache)) % 1_000_000_007
	res = res % 1_000_000_007
	cache[startRow][startColumn][maxMove] = &res
	return res
}
func main() {
	fmt.Println(390153306 == findPaths(36, 5, 50, 15, 3))
	fmt.Println(390153306 == findPaths2(36, 5, 50, 15, 3))
}
