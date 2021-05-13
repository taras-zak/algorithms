package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		n := len(res)
		for j := 0; j < n; j++ {
			res = append(res, insertInEveryPosition(res[j], nums[i])...)
		}
		res = res[n:]

	}
	return res
}

func insertInEveryPosition(arr []int, i int) [][]int {
	if len(arr) == 0 {
		return append([][]int(nil), append(arr, i))
	}
	var res [][]int
	for pos := 0; pos <= len(arr); pos++ {
		tmp := append([]int(nil), arr...)
		tmp = append(tmp[:pos], append([]int{i}, tmp[pos:]...)...)
		res = append(res, tmp)
	}
	return res
}

func permuteRecursive(nums []int) [][]int {
	return backtrackPermute([][]int{}, 0, len(nums), nums)
}

func backtrackPermute(arr [][]int, start, end int, nums []int) [][]int {
	if start == end {
		arr = append(arr, append([]int(nil), nums...))
	}
	for i := start; i < end; i++ {
		nums[start], nums[i] = nums[i], nums[start]
		arr = backtrackPermute(arr, start+1, end, nums)
		nums[start], nums[i] = nums[i], nums[start]
	}
	return arr
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permuteRecursive([]int{1, 2, 3}))
}
