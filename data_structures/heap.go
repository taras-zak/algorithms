package main

type IntHeap struct {
	arr []int
}

func NewHeap(a []int) *IntHeap {
	h := &IntHeap{arr: a}
	// heapify
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		h.downHeap(i, len(h.arr))
	}
	return &IntHeap{arr: a}
}

func (h *IntHeap) Len() int {
	return len(h.arr)
}

func (h *IntHeap) Insert(el int) {
	h.arr = append(h.arr, el)
	i := len(h.arr) - 1
	parent := (i - 1) / 2
	for parent >= 0 && h.arr[parent] > h.arr[i] {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
		parent = (i - 1) / 2
	}
}

func (h *IntHeap) downHeap(i int, end int) {
	for {
		smallest := i
		left := i*2 + 1
		right := i*2 + 2
		if left < end && h.arr[smallest] > h.arr[left] {
			smallest = left
		}
		if right < end && h.arr[smallest] > h.arr[right] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}

func (h *IntHeap) Pop() int {
	res := h.arr[0]
	last := len(h.arr) - 1
	h.arr[0] = h.arr[last]
	h.arr = h.arr[:last]
	h.downHeap(0, len(h.arr))
	return res
}

func (h *IntHeap) Min() int {
	return h.arr[0]
}

// HeapSort make sort inplace in descending order
func (h *IntHeap) HeapSort() {
	end := len(h.arr)
	for i := len(h.arr) - 1; i > 0; i-- {
		h.arr[0], h.arr[i] = h.arr[i], h.arr[0]
		end--
		h.downHeap(0, end)
	}
}
