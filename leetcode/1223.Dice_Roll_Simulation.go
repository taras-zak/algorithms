package main

import "fmt"

const mod = 1e9 + 7

func dieSimulator(n int, rollMax []int) int {
	dp := [5001][7]int{}
	dp[0][0] = 1
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	dp[1][0] = 6
	for roll := 2; roll <= n; roll++ {
		for num := 1; num <= 6; num++ {
			for repeat := 1; repeat <= rollMax[num-1]; repeat++ {
				if repeat > roll {
					break
				}
				dp[roll][num] = addMod(dp[roll][num], dp[roll-repeat][0]-dp[roll-repeat][num])
			}
		}
		dp[roll][0] = sum(dp[roll])
	}
	return dp[n][0] % mod
}

func addMod(a, b int) int {
	return (a + b) % mod
}

func sum(arr [7]int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
}
