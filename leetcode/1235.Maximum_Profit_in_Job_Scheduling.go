package main

import (
	"fmt"
	"sort"
)

type node struct {
	start, end, profit int
}

// O(n^2)
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]*node, 0, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, &node{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end <= jobs[j].end
	})
	dp := make([]int, len(jobs))
	dp[0] = jobs[0].profit
	for i := 1; i < len(jobs); i++ {
		dp[i] = max(dp[i-1], jobs[i].profit)
		for j := i - 1; j >= 0; j-- {
			if jobs[i].start >= jobs[j].end {
				dp[i] = max(dp[i], dp[j]+jobs[i].profit)
				break
			}
		}
	}
	maxProfit := dp[0]
	for _, i := range dp {
		maxProfit = max(maxProfit, i)
	}
	return maxProfit
}

// O(n log(n))
func jobScheduling2(startTime []int, endTime []int, profit []int) int {
	jobs := make([]*node, 0, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, &node{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	dp := make([]int, len(jobs))
	dp[0] = jobs[0].profit
	for i := 1; i < len(jobs); i++ {
		p := jobs[i].profit
		l := search(jobs, i)
		if l != -1 {
			p += dp[l]
		}
		dp[i] = max(dp[i-1], p)
	}
	return dp[len(jobs)-1]
}

func search(jobs []*node, index int) int {
	start, end := 0, index-1
	for start <= end {
		mid := (start + end) / 2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				start = mid + 1
			} else {
				return mid
			}
		} else {
			end = mid - 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	start := []int{1, 2, 3, 4, 6}
	end := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	fmt.Println(jobScheduling2(start, end, profit))
}
