package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		n := len(res)
		for j := 0; j < n; j++ {
			res = append(res, append(append([]int(nil), res[j]...), nums[i]))
		}
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 2}))
}
