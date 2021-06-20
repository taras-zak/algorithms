package main

import "fmt"

func mergeStones(stones []int, k int) int {
	if (len(stones)-1)%(k-1) != 0 {
		return -1
	}

	prefixSum := make([]int, len(stones)+1)
	for i := 1; i <= len(stones); i++ {
		prefixSum[i] = prefixSum[i-1] + stones[i-1]
	}

	var dp [30][30]int
	for size := k; size <= len(stones); size++ {
		for left := 0; left < len(stones)-size+1; left++ {
			right := left + size - 1
			dp[left][right] = 1<<31 - 1
			for mid := left; mid < right; mid += k - 1 {
				dp[left][right] = min(dp[left][right], dp[left][mid]+dp[mid+1][right])
			}
			if (size-1)%(k-1) == 0 {
				dp[left][right] += prefixSum[right+1] - prefixSum[left]
			}
		}
	}
	return dp[0][len(stones)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2))
}
