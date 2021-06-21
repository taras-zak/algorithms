package main

import "fmt"

func getMoneyAmountTD(n int) int {
	var dp [200][200]int
	return dfs(1, n, dp)
}

func dfs(l, r int, dp [200][200]int) int {
	if l >= r {
		return 0
	}
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	res := 1<<31 - 1
	for x := l; x <= r; x++ {
		res = min(res, x+max(dfs(l, x-1, dp), dfs(x+1, r, dp)))
	}
	dp[l][r] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMoneyAmount(n int) int {
	var dp [201][201]int
	for size := 2; size <= n; size++ {
		for left := 1; left <= n-size+1; left++ {
			right := left + size - 1
			res := 1<<31 - 1
			for k := left + (size-1)/2; k < right; k++ {
				res = min(res, k+max(dp[left][k-1], dp[k+1][right]))
			}
			dp[left][right] = res
		}
	}
	return dp[1][n]
}

func main() {
	fmt.Println(getMoneyAmountTD(10))
	fmt.Println(getMoneyAmount(10))
}
