package main

import "fmt"

func numRollsToTarget(d int, f int, target int) int {
	mod := 1_000_000_007
	dp := [31][1001]int{}
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 1; j <= f; j++ {
			for t := j; t <= target; t++ {
				dp[i][t] = (dp[i][t] + dp[i-1][t-j]) % mod
			}
		}
	}
	return dp[d][target]
}

func main() {
	fmt.Println(numRollsToTarget(2, 6, 7))
}
