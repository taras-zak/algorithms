package main

import (
	"fmt"
	"strconv"
)

func nextGreaterElement3(n int) int {
	nums := make([]int, 0)
	for _, char := range strconv.Itoa(n) {
		nums = append(nums, int(char-'0'))
	}

	var k = -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		return -1
	}
	var i int = 0
	for j := k + 1; j < len(nums); j++ {
		if nums[k] < nums[j] {
			i = j
		}
	}
	nums[k], nums[i] = nums[i], nums[k]
	for a, b := k+1, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
	res := arrToInt(nums)
	if res > (1<<31)-1 {
		return -1
	}
	return res
}

func arrToInt(nums []int) int {
	res := 0
	op := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res += nums[i] * op
		op *= 10
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement3(234))
}
