package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	dayIndex := 0
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)

	for i := 1; i <= lastDay; i++ {
		if i != days[dayIndex] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = min(dp[i-1]+costs[0], dp[max(0, i-7)]+costs[1], dp[max(0, i-30)]+costs[2])
			dayIndex++
		}
	}
	return dp[lastDay]
}

func min(a ...int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))
}
