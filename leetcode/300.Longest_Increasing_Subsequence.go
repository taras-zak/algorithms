package main

import "fmt"

// patience sorting
func lengthOfLIS(nums []int) int {
	pile := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		i, j := 0, size
		for i < j {
			m := (j + i) / 2
			if pile[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		pile[i] = num
		if i == size {
			size++
		}
	}
	return size
}

func lengthOfLISDp(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[i] <= dp[j] {
				dp[i] = 1 + dp[j]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISDp([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
