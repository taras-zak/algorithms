package main

import "fmt"

func combinationSum4(candidates []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for sum := 1; sum <= target; sum++ {
		for _, num := range candidates {
			if sum-num >= 0 {
				dp[sum] += dp[sum-num]
			}

		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
}
