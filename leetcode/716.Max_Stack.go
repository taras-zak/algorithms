package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type entry struct {
	Value   int
	HeapIdx int
}

type MaxHeap []*list.Element

func (m MaxHeap) Len() int { return len(m) }
func (m MaxHeap) Less(a, b int) bool {
	return m[a].Value.(*entry).Value >= m[b].Value.(*entry).Value
}
func (m MaxHeap) Swap(a, b int) {
	temp := m[a]
	m[a] = m[b]
	m[a].Value.(*entry).HeapIdx = a
	m[b] = temp
	m[b].Value.(*entry).HeapIdx = b
}

func (m *MaxHeap) Push(a interface{}) {
	element := a.(*list.Element)
	element.Value.(*entry).HeapIdx = len(*m)
	*m = append(*m, element)
}

func (m *MaxHeap) Pop() interface{} {
	last := len(*m) - 1
	top := (*m)[last]
	*m = (*m)[:last]
	return top
}

type MaxStack struct {
	maxHeap *MaxHeap
	list    *list.List
}

func Constructor() *MaxStack {
	return &MaxStack{
		maxHeap: &MaxHeap{},
		list:    list.New(),
	}
}

func (ms *MaxStack) Push(x int) {
	ms.list.PushFront(&entry{Value: x})
	element := ms.list.Front()
	heap.Push(ms.maxHeap, element)
	return
}

func (ms *MaxStack) Pop() int {
	if ms.list.Len() == 0 {
		return -1
	}
	top := ms.list.Front()
	ms.list.Remove(top)
	heap.Remove(ms.maxHeap, top.Value.(*entry).HeapIdx)
	return top.Value.(*entry).Value
}

func (ms *MaxStack) Top() int {
	if ms.list.Len() == 0 {
		return -1
	}
	return ms.list.Front().Value.(*entry).Value
}

func (ms *MaxStack) PeekMax() int {
	if ms.list.Len() == 0 {
		return -1
	}
	return (*ms.maxHeap)[0].Value.(*entry).Value
}

func (ms *MaxStack) PopMax() int {
	maxElement := heap.Pop(ms.maxHeap).(*list.Element)
	ms.list.Remove(maxElement)
	return maxElement.Value.(*entry).Value
}

func main() {
	stack := Constructor()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	fmt.Println(stack.Top())
	fmt.Println(stack.PeekMax())
	fmt.Println(stack.Pop())
	fmt.Println(stack.PopMax())

	fmt.Println(stack.Top())
	fmt.Println(stack.PeekMax())
}

[
[1,1,1,1,1],
[1,0,0,0,1],
[1,0,1,0,1],
[1,0,0,0,1],
[1,1,1,1,1]]