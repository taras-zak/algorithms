package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	// var dp [1001][1001]int
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(s)+1)
		dp[i][i] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func longestPalindromeSubseqOpt(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		prev := 0
		for j := i + 1; j < len(s); j++ {
			tmp := dp[j]
			if s[i] == s[j] {
				dp[j] = 2 + prev
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}
	return dp[len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseqOpt("bbbab"))
}
