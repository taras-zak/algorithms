package main

import "fmt"

func findLengthOfLCISTwoPointer(nums []int) int {
	l, r := 0, 1
	max := 1
	for r < len(nums) {
		if nums[r-1] < nums[r] {
			if max < r-l+1 {
				max = r - l + 1
			}
		} else if nums[r-1] == nums[r] {
			l = r
			if max < 1 {
				max = 1
			}
		} else {
			l = r
		}
		r++
	}
	return max
}

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = 1 + dp[i-1]
		} else {
			dp[i] = 1
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCISTwoPointer([]int{1, 3, 5, 4, 7}))
}
