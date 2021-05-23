package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	dp := make([][][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([][2]int, 0)
	}
	size := 0
	for _, num := range nums {
		i, j := 0, size
		for i < j {
			m := (j + i) / 2
			if dp[m][len(dp[m])-1][0] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		opt := 1
		row := i - 1
		if row >= 0 {
			l, r := 0, len(dp[row])
			for l < r {
				m := (l + r) / 2
				if dp[row][m][0] < num {
					r = m
				} else {
					l = m + 1
				}
			}
			opt = dp[row][len(dp[row])-1][1]
			if l != 0 {
				opt -= dp[row][l-1][1]
			}
		}

		pileLen := len(dp[i])
		if pileLen != 0 {
			opt += dp[i][pileLen-1][1]
		}
		dp[i] = append(dp[i], [2]int{num, opt})
		if i == size {
			size++
		}
	}
	return dp[size-1][len(dp[size-1])-1][1]
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
}
