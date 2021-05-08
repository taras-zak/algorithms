package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	return res
}

func main() {
	fmt.Println(myPow(2, -2))
}
