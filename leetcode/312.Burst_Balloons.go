package main

import "fmt"

func maxCoins(nums []int) int {
	values := append([]int{1}, nums...)
	values = append(values, 1)
	var dp [502][502]int
	for size := 2; size < len(values); size++ {
		for left := 0; left < len(values)-size; left++ {
			right := left + size
			for k := left + 1; k < right; k++ {
				dp[left][right] = max(dp[left][right], values[left]*values[k]*values[right]+dp[left][k]+dp[k][right])
			}
		}
	}
	return dp[0][len(values)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
