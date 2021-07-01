package main

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	h := NewHeap([]int{})
	h.Insert(5)
	h.Insert(2)
	h.Insert(9)
	h.Insert(3)
	if h.Min() != 2 {
		t.Errorf("expected %d, got %d", 2, h.Min())
	}
}

func TestHeapPop(t *testing.T) {
	h := NewHeap([]int{1, 2})
	if x := h.Pop(); x != 1 {
		t.Errorf("expected pop %d, got %d", 1, x)
	}
	if h.Len() != 1 {
		t.Errorf("expected len %d, got %d", 1, h.Len())
	}
	if x := h.Pop(); x != 2 {
		t.Errorf("expected pop %d, got %d", 2, x)
	}
	if h.Len() != 0 {
		t.Errorf("expected len %d, got %d", 0, h.Len())
	}
}

func TestHeapNewHeap(t *testing.T) {
	h := NewHeap([]int{5, 4, 3, 2, 1})
	res := make([]int, 0)
	for h.Len() != 0 {
		res = append(res, h.Pop())
	}
	for i := 0; i < 5; i++ {
		if res[i] != i+1 {
			t.Error("expected sorted res")
		}
	}
}

func TestHeapHeapSort(t *testing.T) {
	h := NewHeap([]int{3, 1, 2, 5, 4})
	h.HeapSort()
	exp := 1
	for i := 4; i >= 0; i-- {
		if h.arr[i] != exp {
			t.Error("expected sorted res desc")
		}
		exp++
	}
}
