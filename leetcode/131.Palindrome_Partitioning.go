package main

import "fmt"

func partition(s string) [][]string {
	res := make([][]string, 0)
	res, _ = backtrack(s, 0, res, []string{})
	return res
}

func backtrack(s string, start int, res [][]string, curr []string) ([][]string, []string) {
	if start >= len(s) {
		res = append(res, append([]string{}, curr...))
		return res, curr
	}
	for i := start; i < len(s); i++ {
		if isPal(s, start, i) {
			curr = append(curr, s[start:i+1])
			res, curr = backtrack(s, i+1, res, curr)
			curr = curr[:len(curr)-1]
		}
	}
	return res, curr
}

func isPal(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
