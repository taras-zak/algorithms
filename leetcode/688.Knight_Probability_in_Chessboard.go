package main

import "fmt"

func knightProbability(n int, maxMove int, startRow int, startColumn int) float64 {
	var cache [25][25][101]*float64
	return dfsKnightProbability(n, maxMove, startRow, startColumn, &cache)
}

var dirs = [][]int{
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	{-2, 1}, {-1, 2}, {1, 2}, {2, 1},
}

func dfsKnightProbability(n, maxMove, startRow, startColumn int, cache *[25][25][101]*float64) float64 {
	if startRow < 0 || startRow >= n || startColumn < 0 || startColumn >= n {
		return 0
	}
	if maxMove == 0 {
		return 1
	}
	if cache[startRow][startColumn][maxMove] != nil {
		return *cache[startRow][startColumn][maxMove]
	}
	var res float64
	for _, dir := range dirs {
		res += 0.125 * dfsKnightProbability(n, maxMove-1, startRow+dir[0], startColumn+dir[1], cache)
	}
	cache[startRow][startColumn][maxMove] = &res
	return res
}

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
}
