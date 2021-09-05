package main

import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	sb := []rune(s)
	if k > 1 {
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})
		return string(sb)
	}
	res := sb
	for i := 0; i < len(sb); i++ {
		sb = append(sb[1:], sb[0:1]...)
		res = min(res, sb)
	}
	return string(res)
}

func min(a, b []rune) []rune {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			return b
		} else {
			return a
		}
	}
	return a
}

func main() {
	fmt.Println(orderlyQueue("cba", 1))
	fmt.Println(orderlyQueue("baaca", 3))
}
