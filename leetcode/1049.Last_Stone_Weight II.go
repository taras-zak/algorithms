package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}

	dp := make([]int, (sum/2)+1)

	for n := 0; n < len(stones); n++ {
		for s := sum / 2; s >= stones[n]; s-- {
			dp[s] = max(dp[s], dp[s-stones[n]]+stones[n])
		}
	}
	return sum - 2*dp[sum/2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
