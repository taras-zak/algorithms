package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		ones, zeros := count(str)
		for z := m; z >= zeros; z-- {
			for o := n; o >= ones; o-- {
				dp[z][o] = max(dp[z][o], dp[z-zeros][o-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(str string) (int, int) {
	var ones, zeros int
	for _, ch := range str {
		if ch == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return ones, zeros
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
