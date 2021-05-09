package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for step := 2; step <= len(s); step++ {
		if s[step-1] != '0' {
			dp[step] += dp[step-1]
		}
		d, _ := strconv.Atoi(s[step-2 : step])
		if d >= 10 && d <= 26 {
			dp[step] += dp[step-2]
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(numDecodings("226"))
}
