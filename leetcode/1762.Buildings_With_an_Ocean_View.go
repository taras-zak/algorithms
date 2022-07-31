package main

import (
	"fmt"
)

func findBuildings(heights []int) []int {
	stack := []int{}
	var res []int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top < heights[i] {
				stack = stack[:len(stack)-1]
				res = res[:len(res)-1]
			} else {
				break
			}
		}
		stack = append(stack, heights[i])
		res = append(res, i)
	}
	return res
}

func findBuildingsOpt(heights []int) []int {
	var res []int
	prevHight := 0
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > prevHight {
			res = append(res, i)
			prevHight = heights[i]
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func main() {
	arr := []int{4, 2, 3, 1}
	fmt.Println(findBuildings(arr))
}
