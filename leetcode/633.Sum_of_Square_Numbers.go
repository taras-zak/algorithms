package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	if c == 1 || c == 0 {
		return true
	}
	l, r := 1, int(math.Sqrt(float64(c)))
	for l <= r {
		res := l*l + r*r
		if res == c {
			return true
		}
		if res > c {
			r--
		} else {
			l--
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(25))
}
