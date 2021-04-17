package main

import "github.com/taras-zak/algorithms/utils"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for a := 1; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			if coins[c] <= a {
				dp[a] = utils.Min(dp[a], dp[a-coins[c]]+1)
			}
		}
	}
	if dp[len(dp)-1] == amount+1 {
		return -1
	}
	return dp[len(dp)-1]
}

func main() {
	println(coinChange([]int{1, 2, 3}, 5))
	println(coinChange([]int{1, 2, 3}, 11))
	println(coinChange([]int{1, 2, 3}, 7))
	println(coinChange([]int{3, 5}, 7))
}
