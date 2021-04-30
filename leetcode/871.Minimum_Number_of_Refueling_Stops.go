package main

import (
	"container/heap"
	"fmt"
)

func minRefuelStopsDp(target int, startFuel int, stations [][]int) int {
	dp := make([]int, len(stations)+1)
	dp[0] = startFuel
	for i := 0; i < len(stations); i++ {
		for t := i; t >= 0; t-- {
			if dp[t] >= stations[i][0] {
				dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}
	for t := 0; t <= len(stations); t++ {
		if dp[t] >= target {
			return t
		}
	}
	return -1
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	pq := make(MaxHeap, 0)
	var i, res int
	curr := startFuel
	for curr < target {
		for i < len(stations) && curr >= stations[i][0] {
			heap.Push(&pq, stations[i][1])
			i++
		}
		if len(pq) == 0 {
			return -1
		}
		curr += heap.Pop(&pq).(int)
		res++
	}
	return res
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func (h *MaxHeap) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func main() {
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
	fmt.Println(minRefuelStopsDp(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}
