package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := []int{0}
	for len(dp) <= n {
		i := len(dp)
		minCount := math.MaxInt64
		for j := 1; j*j <= i; j++ {
			minCount = min(minCount, dp[i-j*j]+1)
		}
		dp = append(dp, minCount)
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Println(i, numSquares(i))
	}
}
