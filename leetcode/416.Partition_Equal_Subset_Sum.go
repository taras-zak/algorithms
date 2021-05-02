package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	dp := make([]bool, sum+1)
	dp[0] = true

	for n := 0; n < len(nums); n++ {
		for s := sum; s > 0; s-- {
			if s >= nums[n] {
				dp[s] = dp[s] || dp[s-nums[n]]
			}
		}
	}
	return dp[sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
