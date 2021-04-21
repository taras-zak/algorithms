package main

import (
	"container/heap"
	"fmt"
)

const MAX = 10000

type item struct {
	val, pos int
}

type MaxHeap []item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minFallingPathSum2Heap(arr [][]int) int {
	n := len(arr)
	h := &MaxHeap{}
	for i := 0; i < n; i++ {
		heap.Push(h, item{arr[0][i], i})
		if h.Len() > 2 {
			heap.Pop(h)
		}
	}
	for i := 1; i < n; i++ {
		minSecond, minFirst := heap.Pop(h).(item), heap.Pop(h).(item)
		h = &MaxHeap{}
		for j := 0; j < n; j++ {
			if minFirst.pos != j {
				arr[i][j] += minFirst.val
			} else {
				arr[i][j] += minSecond.val
			}
			heap.Push(h, item{arr[i][j], j})
			if h.Len() > 2 {
				heap.Pop(h)
			}
		}
	}
	min := arr[n-1][0]
	for i := 1; i < n; i++ {
		if min > arr[n-1][i] {
			min = arr[n-1][i]
		}
	}
	return min
}

func minFallingPathSum2(arr [][]int) int {
	n := len(arr)
	for i := 1; i < n; i++ {
		minFirst, minSecond := minTwo(arr[i-1])
		for j := 0; j < n; j++ {
			if arr[i-1][j] == minFirst {
				arr[i][j] += minSecond
			} else {
				arr[i][j] += minFirst
			}
		}
	}
	min := arr[n-1][0]
	for i := 1; i < n; i++ {
		if min > arr[n-1][i] {
			min = arr[n-1][i]
		}
	}
	return min
}

func minTwo(a []int) (int, int) {
	minF, minS := MAX, MAX
	for _, val := range a {
		if val < minF {
			minS = minF
			minF = val
		} else if val < minS {
			minS = val
		}
	}
	return minF, minS
}

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum2Heap(m))
	m = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum2(m))
}
