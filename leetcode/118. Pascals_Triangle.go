package main

import (
	"fmt"
)

func generate(n int) [][]int {
	res := [][]int{{1}}
	for i := 1; i < n; i++ {
		var level []int
		for j := 0; j <= i; j++ {
			sum := 0
			if j-1 >= 0 {
				sum += res[i-1][j-1]
			}
			if j < len(res[i-1]) {
				sum += res[i-1][j]
			}
			level = append(level, sum)
		}
		res = append(res, level)
	}
	return res
}

func main() {
	fmt.Println(generate(10))
}
