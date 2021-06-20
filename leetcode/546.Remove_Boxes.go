package main

import "fmt"

func removeBoxesRec(boxes []int) int {
	var dp [100][100][100]int
	return dfs(boxes, 0, len(boxes)-1, 0, dp)
}

func dfs(boxes []int, left, right, k int, dp [100][100][100]int) int {
	if left > right {
		return 0
	}
	if dp[left][right][k] != 0 {
		return dp[left][right][k]
	}
	for left < right && boxes[left] == boxes[left+1] {
		left++
		k++
	}
	res := (k+1)*(k+1) + dfs(boxes, left+1, right, 0, dp)
	for m := left + 1; m <= right; m++ {
		if boxes[left] == boxes[m] {
			res = max(res, dfs(boxes, left+1, m-1, 0, dp)+dfs(boxes, m, right, k+1, dp))
		}
	}
	dp[left][right][k] = res
	return res
}

func removeBoxes(boxes []int) int {
	var dp [101][101][101]int

	for i := 0; i < len(boxes); i++ {
		for k := 0; k <= i; k++ {
			dp[i][i][k] = (k + 1) * (k + 1)
		}
	}

	for size := 1; size <= len(boxes); size++ {
		for left := 0; left < len(boxes)-size+1; left++ {
			right := left + size - 1

			for k := 0; k <= left; k++ {
				res := (k+1)*(k+1) + dp[left+1][right][0]
				for m := left + 1; m <= right; m++ {
					if boxes[left] == boxes[m] {
						res = max(res, dp[left+1][m-1][0]+dp[m][right][k+1])
					}
				}
				dp[left][right][k] = res
			}
		}
	}
	return dp[0][len(boxes)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(removeBoxesRec([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}
