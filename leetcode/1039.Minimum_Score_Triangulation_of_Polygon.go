package main

import "fmt"

// top-down
func minScoreTriangulation(values []int) int {
	return dfs(0, len(values)-1, values)
}

func dfs(l, r int, v []int) int {
	if r-l+1 < 3 {
		return 0
	}
	m := 1<<31 - 1
	for k := l + 1; k < r; k++ {
		m = min(m, v[l]*v[k]*v[r]+dfs(l, k, v)+dfs(k, r, v))
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minScoreTriangulationDp(values []int) int {
	var dp [50][50]int
	for size := 2; size < len(values); size++ {
		for left := 0; left < len(values)-size; left++ {
			right := left + size
			dp[left][right] = 1<<31 - 1
			for k := left + 1; k < right; k++ {
				dp[left][right] = min(dp[left][right], values[left]*values[right]*values[k]+dp[left][k]+dp[k][right])
			}
		}
	}
	return dp[0][len(values)-1]
}

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(minScoreTriangulationDp([]int{1, 2, 3, 4, 5, 6}))
}
