package main

import "fmt"

func nextGreaterElementNaive(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = -1
	}
	for i, num1 := range nums1 {
		i2 := -1
		for j, num2 := range nums2 {
			if num2 == num1 {
				i2 = j
			} else if num2 > num1 && i2 != -1 {
				res[i] = num2
				break
			}
		}
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	stack := make([]int, 0, len(nums2))
	for _, num := range nums2 {
		var top int
		for len(stack) > 0 && stack[len(stack)-1] < num {
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res[top] = num
		}
		stack = append(stack, num)
	}
	for i := 0; i < len(nums1); i++ {
		if n, ok := res[nums1[i]]; ok {
			nums1[i] = n
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}

func main() {
	fmt.Println(nextGreaterElementNaive([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
