package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	return backtrackCombineSum3([][]int{}, []int{}, 1, n, k)
}

func backtrackCombineSum3(res [][]int, temp []int, start, target, k int) [][]int {
	if target == 0 && k == 0 {
		res = append(res, append([]int(nil), temp...))
		return res
	}
	for i := start; i < 10 && target >= i; i++ {
		temp = append(temp, i)
		res = backtrackCombineSum3(res, temp, i+1, target-i, k-1)
		temp = temp[:len(temp)-1]
	}
	return res
}

func main() {
	fmt.Println(combinationSum3(9, 45))
}
