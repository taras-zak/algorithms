package main

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, 0)
	temp := make([]int, 0)
	pivotCount := 0
	for _, num := range nums {
		if num == pivot {
			pivotCount++
			continue
		}
		if num <= pivot {
			res = append(res, num)
		} else {
			temp = append(temp, num)
		}
	}
	for pivotCount != 0 {
		res = append(res, pivot)
		pivotCount--
	}
	return append(res, temp...)
}
