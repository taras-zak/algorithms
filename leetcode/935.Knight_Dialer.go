package main

import "fmt"

func knightDialer(n int) int {
	move := [10][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{8, 4},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{6, 2},
		{1, 3},
		{4, 2},
	}
	dp := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		dpCurr := [10]int{}
		for j := 0; j < 10; j++ {
			for _, next := range move[j] {
				dpCurr[j] = (dpCurr[j] + dp[next]) % 1_000_000_007
			}
		}
		dp = dpCurr

	}
	res := 0
	for i := 0; i < 10; i++ {
		res = (res + dp[i]) % 1_000_000_007
	}
	return res
}

func main() {
	fmt.Println(knightDialer(4))
}
