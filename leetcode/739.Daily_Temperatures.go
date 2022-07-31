package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([][2]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top[0] <= temperatures[i] {
				stack = stack[:len(stack)-1]
			} else {
				res[i] = top[1] - i
				stack = append(stack, [2]int{temperatures[i], i})
				break
			}
		}
		if len(stack) == 0 {
			stack = append(stack, [2]int{temperatures[i], i})
		}
	}
	return res
}

func main() {
	arr := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	fmt.Println(dailyTemperatures(arr))
}
