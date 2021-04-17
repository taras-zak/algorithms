package main

import (
	"fmt"
	"github.com/taras-zak/algorithms/utils"
	"math"
)

func numSquares(n int) int {
	dp := []int{0}
	for len(dp) <= n {
		i := len(dp)
		minCount := math.MaxInt64
		for j := 1; j*j <= i; j++ {
			minCount = utils.Min(minCount, dp[i-j*j]+1)
		}
		dp = append(dp, minCount)
	}
	return dp[n]
}

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Println(i, numSquares(i))
	}
}
