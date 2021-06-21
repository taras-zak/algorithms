package main

import "fmt"

func guess(n int) int {
	if n == 64 {
		return 0
	} else if n < 64 {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	l := 1
	r := n
	for l < r {
		mid := l + (r-l)/2
		g := guess(mid)
		if g > 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	fmt.Println(guessNumber(100))
}
