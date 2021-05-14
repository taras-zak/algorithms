package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return backtrackCombineSum([][]int{}, []int{}, candidates, 0, target)
}

func backtrackCombineSum(res [][]int, temp []int, candidates []int, start, target int) [][]int {
	if target == 0 {
		res = append(res, append([]int(nil), temp...))
		return res
	}

	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		temp = append(temp, candidates[i])
		res = backtrackCombineSum(res, temp, candidates, i, target-candidates[i])
		temp = temp[:len(temp)-1]
	}
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
