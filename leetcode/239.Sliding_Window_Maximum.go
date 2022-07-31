package main

import (
	"container/list"
	"fmt"
)

type Item struct {
	Index int
	Value int
}

func maxSlidingWindow(nums []int, k int) []int {
	l := list.New()
	var ans []int

	for i, val := range nums {
		if l.Len() > 0 && l.Front().Value.(*Item).Index <= i-k {
			l.Remove(l.Front())
		}

		for l.Len() > 0 && l.Back().Value.(*Item).Value < val {
			l.Remove(l.Back())
		}

		l.PushBack(&Item{Index: i, Value: val})

		if i >= k-1 {
			ans = append(ans, l.Front().Value.(*Item).Value)
		}
	}
	return ans
}

func maxSlidingWindowOptimized(nums []int, k int) []int {
	var deq []int
	var ans []int

	for i := range nums {
		for len(deq) > 0 && deq[0] <= i-k {
			deq = deq[1:]
		}

		for len(deq) > 0 && nums[deq[len(deq)-1]] < nums[i] {
			deq = deq[:len(deq)-1]
		}

		deq = append(deq, i)

		if i >= k-1 {
			ans = append(ans, nums[deq[0]])
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindowOptimized([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println([]int{3, 3, 5, 5, 6, 7})
}
