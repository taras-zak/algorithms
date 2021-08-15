package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	start, end := 0, 0
	minLength := 1000000
	minLengthStart := 0
	tMap := make(map[byte]int)
	for _, r := range []byte(t) {
		tMap[r]++
	}
	counter := len(t)
	for end < len(s) {
		if tMap[s[end]] > 0 {
			counter--
		}
		tMap[s[end]]--
		end++
		for counter == 0 {
			if minLength > end-start {
				minLengthStart = start
				minLength = end - start
			}
			tMap[s[start]]++
			if tMap[s[start]] == 1 {
				counter++
			}
			start++
		}
	}
	if minLength == 1000000 {
		return ""
	}
	return s[minLengthStart : minLengthStart+minLength]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
