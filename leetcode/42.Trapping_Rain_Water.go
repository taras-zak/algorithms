package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	leftMax := height[0]
	rightMax := height[n-1]
	left := 1
	right := n - 2
	sum := 0
	for left <= right {
		if leftMax < rightMax {
			if height[left] < leftMax {
				sum += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				sum += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return sum
}

func trapOpt(height []int) int {
	maxLeft := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		if height[i-1] > maxLeft[i-1] {
			maxLeft[i] = height[i-1]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}
	maxRight := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		if height[i+1] > maxRight[i+1] {
			maxRight[i] = height[i+1]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	var sum int
	for i := 1; i < len(height)-1; i++ {
		minLevel := maxLeft[i]
		if minLevel > maxRight[i] {
			minLevel = maxRight[i]
		}
		level := minLevel - height[i]
		if level > 0 {
			sum += level
		}
	}
	return sum
}

func trapNaive(height []int) int {
	var sum int
	for i := 1; i < len(height)-1; i++ {
		maxLeft := max(height[:i])
		maxRight := max(height[i+1:])
		minLevel := maxLeft
		if maxLeft > maxRight {
			minLevel = maxRight
		}
		level := minLevel - height[i]
		if level > 0 {
			sum += level
		}
	}
	return sum
}

func max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	//arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	arr := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trapNaive(arr))
	fmt.Println(trap(arr))
}
