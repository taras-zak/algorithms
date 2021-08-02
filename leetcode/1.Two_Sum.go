package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		} else {
			m[num] = i
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
