package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += check(s, i, i)
		res += check(s, i, i+1)
	}
	return res
}

func check(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		res++
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("abccbz"))
}
