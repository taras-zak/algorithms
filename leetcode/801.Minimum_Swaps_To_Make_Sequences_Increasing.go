package main

import "fmt"

func minSwap(nums1 []int, nums2 []int) int {
	var INF = len(nums1) + 1
	lastSwapedMin := 1
	lastNotSwapedMin := 0
	for i := 1; i < len(nums1); i++ {
		currSwappedMin := INF
		currNotSwappedMin := INF
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			// cur swapped, last swaped
			currSwappedMin = min(lastSwapedMin+1, currSwappedMin)
			// cur not swapped, last not swaped
			currNotSwappedMin = min(lastNotSwapedMin, currNotSwappedMin)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			// cur swapped, last not swaped
			currSwappedMin = min(lastNotSwapedMin+1, currSwappedMin)
			// cur not swapped, last swaped
			currNotSwappedMin = min(lastSwapedMin, currNotSwappedMin)
		}
		lastSwapedMin, lastNotSwapedMin = currSwappedMin, currNotSwappedMin
	}
	return min(lastSwapedMin, lastNotSwapedMin)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
}
