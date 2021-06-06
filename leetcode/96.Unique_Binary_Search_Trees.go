package main

import "fmt"

// Catalan numbers
func numTreesCatalan(n int) int {
	curr := 1
	for i := 1; i < n; i++ {
		curr = (2 * (2*i + 1) * curr) / (i + 2)
	}
	return curr
}

// DP
func numTrees(n int) int {
	dp := [20]int{}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// j – number in root
		for j := 1; j <= i; j++ {
			// dp[j-1] – left tree, dp[i-j] – right tree
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(5))
	fmt.Println(numTreesCatalan(5))
}
