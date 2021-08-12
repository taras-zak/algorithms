package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, str := range strs {
		hash := hashString(str)
		m[hash] = append(m[hash], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func hashString(str string) [26]byte {
	var counts [26]byte
	for _, r := range []rune(str) {
		counts[r-'a']++
	}
	return counts
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
