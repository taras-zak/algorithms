package main

import "fmt"

func longestPalindromeDp(s string) string {
	var dp [1001][1001]bool
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	palLen := 1
	palStart := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if dp[i+1][j-1] || j-i == 1 {
					dp[i][j] = true
					if palLen < j-i+1 {
						palLen = j - i + 1
						palStart = i
					}
				}
			}
		}
	}
	return s[palStart : palStart+palLen]
}

func longestPalindromeExpand(s string) string {
	palLen := 1
	palStart := 0
	for i := 0; i < len(s); i++ {
		one := expand(s, i, i)
		two := expand(s, i, i+1)
		maxLen := max(one, two)
		if maxLen > palLen {
			palStart = i - (maxLen-1)/2
			palLen = maxLen
		}
	}
	return s[palStart : palStart+palLen]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeDp("yabax"))
	fmt.Println(longestPalindromeExpand("yabax"))
}
