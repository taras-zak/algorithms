package main

import (
	"container/heap"
)

type Heap []*ListNode

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	i := 0
	lastNotNil := 0
	for i < len(lists) {
		if lists[i] != nil {
			lists[lastNotNil] = lists[i]
			lastNotNil++
		}
		i++
	}
	lists = lists[:lastNotNil]

	h := Heap(lists)
	heap.Init(&h)
	head := &ListNode{}
	ret := head
	for h.Len() > 0 {
		top := heap.Pop(&h).(*ListNode)
		if top.Next != nil {
			heap.Push(&h, top.Next)
		}
		head, head.Next = top, top
	}
	return ret.Next
}
