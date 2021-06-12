package main

import "fmt"

func isBadVersion(v int) bool {
	return v > 5
}

func firstBadVersion(n int) int {
	lo, hi := 0, n
	for lo < hi {
		mid := (hi + lo) / 2
		isBad := isBadVersion(mid)
		if isBad {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(firstBadVersion(5))
}
