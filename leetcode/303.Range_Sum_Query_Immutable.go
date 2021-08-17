package main

import "fmt"

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	res := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		res[i+1] = res[i] + nums[i]
	}
	return NumArray{arr: res}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.arr[right+1] - n.arr[left]
}

func main() {
	n := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(n.SumRange(1, 3))
}
