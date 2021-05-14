package main

import "fmt"

func combine(n int, k int) [][]int {
	return backtrackCombine([][]int{}, []int{}, 1, n, k)
}

func backtrackCombine(res [][]int, temp []int, start, n, k int) [][]int {
	if k == 0 {
		res = append(res, append([]int(nil), temp...))
		return res
	}
	for i := start; i <= n-k+1; i++ {
		temp = append(temp, i)
		res = backtrackCombine(res, temp, i+1, n, k-1)
		temp = temp[:len(temp)-1]
	}
	return res
}

func main() {
	fmt.Println(combine(3, 2))
}
