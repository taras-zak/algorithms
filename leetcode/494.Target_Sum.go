package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	totSum := 0
	for i := 0; i < len(nums); i++ {
		totSum += nums[i]
	}
	if totSum < target || (totSum+target)%2 > 0 {
		return 0
	}

	subsetSum := (totSum + target) / 2

	dp := make([]int, subsetSum+1)
	dp[0] = 1
	for n := 0; n < len(nums); n++ {
		for s := subsetSum; s >= nums[n]; s-- {
			dp[s] += dp[s-nums[n]]
		}
	}
	return dp[subsetSum]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 2, 3}, 1))
}
