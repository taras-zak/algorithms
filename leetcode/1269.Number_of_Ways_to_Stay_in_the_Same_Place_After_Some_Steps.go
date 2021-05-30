package main

import "fmt"

func numWays(steps int, arrLen int) int {
	dp := [2][501]int{}
	dp[1][0] = 1
	dp[1][1] = 1
	maxPos := min(steps/2+1, arrLen)
	for step := 2; step <= steps; step++ {
		currStep := step % 2
		prevStep := (step + 1) % 2
		for pos := 0; pos < maxPos; pos++ {
			dp[currStep][pos] = dp[prevStep][pos]
			if pos < maxPos {
				dp[currStep][pos] = addMod(dp[currStep][pos], dp[prevStep][pos+1])
			}
			if pos > 0 {
				dp[currStep][pos] = addMod(dp[currStep][pos], dp[prevStep][pos-1])
			}
		}

	}
	return dp[steps%2][0]
}

func addMod(a, b int) int {
	mod := 1_000_000_007
	return (a + b) % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numWays(3, 2))
}
