package main

import "fmt"

const mod = 1e9 + 7

func countVowelPermutation(n int) int {
	// a, e, i, o, u
	dp := [5]int{1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		dpCurr := [5]int{}
		dpCurr[0] = (dpCurr[0] + dp[1]) % mod
		dpCurr[1] = (dpCurr[1] + dp[0] + dp[2]) % mod
		dpCurr[2] = (dpCurr[2] + dp[0] + dp[1] + dp[3] + dp[4]) % mod
		dpCurr[3] = (dpCurr[3] + dp[2] + dp[4]) % mod
		dpCurr[4] = (dpCurr[4] + dp[0]) % mod
		dp = dpCurr
	}
	res := 0
	for _, i := range dp {
		res = (res + i) % mod
	}
	return res
}

func main() {
	fmt.Println(countVowelPermutation(5))
}
