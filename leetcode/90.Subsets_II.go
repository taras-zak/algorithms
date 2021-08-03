package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	i := 0
	for i < len(nums) {
		count := 0
		for count+i < len(nums) && nums[count+i] == nums[i] {
			count++
		}
		n := len(res)
		for j := 0; j < n; j++ {
			r := append([]int(nil), res[j]...)
			for k := 0; k < count; k++ {
				r = append(r, nums[i])
				res = append(res, r)
			}
		}
		i += count
	}
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
