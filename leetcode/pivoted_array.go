package leetcode

import "github.com/taras-zak/algorithms/algorithms/searching"


func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	pivot := FindPivot(nums, 0, len(nums) - 1)
	if pivot == 0 {
		return searching.BinarySearch(nums, 0, len(nums)-1, target)
	}
	if target < nums[0] {
		return searching.BinarySearch(nums, pivot, len(nums)-1, target)
	} else {
		return searching.BinarySearch(nums, 0, pivot-1, target)
	}
}


func FindPivot(nums []int, low, high int) int {
	if low == high {
		return low
	}
	if nums[0] < nums[len(nums)-1] {
		return 0
	}
	mid := low + (high - low) / 2
	switch {
	case nums[mid] > nums[mid+1]:
		return mid + 1
	case nums[mid] < nums[mid-1]:
		return mid
	case nums[mid] < nums[low]:
		return FindPivot(nums, low, mid-1)
	case nums[high] < nums[mid]:
		return FindPivot(nums, mid+1, high)
	}
	// TODO: rm this
	return mid
}

