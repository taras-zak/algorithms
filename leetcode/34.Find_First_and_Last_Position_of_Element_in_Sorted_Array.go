package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lo := search(nums, target)
	if lo == len(nums) || nums[lo] != target {
		return []int{-1, -1}
	}
	hi := search(nums, target+1) - 1
	return []int{lo, hi}
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	first := len(nums)
	for lo <= hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] >= target {
			first = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return first
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 7, 8, 8, 10}, 7))
}
