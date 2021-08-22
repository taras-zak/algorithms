package main

import "fmt"

func minOperations(s string) int {
	oneStart := 0
	zeroStart := 0
	curr := '1'
	for i, ch := range s {
		if ch == curr {
			zeroStart++
		} else {
			oneStart++
		}
		curr = rune(int('0') + i%2)
	}
	return min(oneStart, zeroStart)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOperations("101110"))
}
