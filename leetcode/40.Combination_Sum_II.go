package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return backtrackCombineSum2([][]int{}, []int{}, candidates, 0, target)
}

func backtrackCombineSum2(res [][]int, temp []int, candidates []int, start, target int) [][]int {
	if target == 0 {
		res = append(res, append([]int(nil), temp...))
		return res
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		temp = append(temp, candidates[i])
		res = backtrackCombineSum2(res, temp, candidates, i+1, target-candidates[i])
		temp = temp[:len(temp)-1]
	}
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
