package main

import "fmt"

func minFlipsMonoIncr(s string) int {
	ones := 0
	flip := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		} else {
			flip++
		}
		flip = min(flip, ones)
	}
	return flip
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFlipsMonoIncr("01011"))
}
