package main

import "fmt"

func mctFromLeafValues(arr []int) int {
	maxInt := 1<<32 - 1
	maxEl := [40][40]int{}
	minSum := [40][40]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			minSum[i][j] = maxInt
		}
		minSum[i][i] = 0
		maxEl[i][i] = arr[i]
	}

	for lenth := 2; lenth <= len(arr); lenth++ {
		for left, right := 0, lenth-1; right < len(arr); left, right = left+1, right+1 {
			for leftPartSize := 1; leftPartSize < lenth; leftPartSize++ {
				rootVal := maxEl[left][left+leftPartSize-1] * maxEl[left+leftPartSize][right]
				sum := minSum[left][left+leftPartSize-1] + minSum[left+leftPartSize][right] + rootVal
				minSum[left][right] = min(minSum[left][right], sum)
			}
			maxEl[left][right] = max(arr[left], maxEl[left+1][right])
		}
	}
	return minSum[0][len(arr)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4, 7, 3, 7, 1, 8, 9}))
}
