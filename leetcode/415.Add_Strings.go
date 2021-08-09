package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	last1 := len(num1) - 1
	last2 := len(num2) - 1
	var flag uint8
	res := make([]uint8, 0)
	for last1 >= 0 || last2 >= 0 || flag == 1 {
		var first uint8
		if last1 >= 0 {
			first = charToInt(num1[last1])
		}
		var second uint8
		if last2 >= 0 {
			second = charToInt(num2[last2])
		}
		dig := (first + second + flag) % 10
		flag = (first + second + flag) / 10
		res = append(res, intToChar(dig))
		last1--
		last2--
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func charToInt(c uint8) uint8 {
	return c - '0'
}

func intToChar(dig uint8) uint8 {
	return '0' + dig
}

func main() {
	fmt.Println(addStrings("7230", "321"))
}
