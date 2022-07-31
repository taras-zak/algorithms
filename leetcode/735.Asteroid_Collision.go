package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
astLoop:
	for i := 0; i < len(asteroids); i++ {
	stackLoop:
		for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroids[i] < 0 {
			if stack[len(stack)-1] > -asteroids[i] {
				break stackLoop
			}
			if stack[len(stack)-1] == -asteroids[i] {
				stack = stack[:len(stack)-1]
				continue astLoop
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, asteroids[i])
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{-2, 1, -1, -2}))
}
