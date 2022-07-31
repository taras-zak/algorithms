package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		prev := res[len(res)-1]
		curr := intervals[i]
		if prev[1] >= curr[1] {
			continue
		}
		if prev[1] >= curr[0] {
			prev[1] = curr[1]
			continue
		}
		res = append(res, curr)
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {5, 10}, {15, 18}}))
}
