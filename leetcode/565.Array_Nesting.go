package main

import "fmt"

func arrayNesting(nums []int) int {
	maxSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}
		size := 0
		for k := i; nums[k] >= 0; size++ {
			numsk := nums[k]
			nums[k] = -1
			k = numsk
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}
