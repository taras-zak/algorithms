package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stack := make([]int, 0, len(nums))
	n := len(nums)
	var num int
	for i := 0; i < 2*n; i++ {
		num = nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{5, 4, 3, 2, 1}))
}
