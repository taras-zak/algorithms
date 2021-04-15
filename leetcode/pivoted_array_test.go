package main

import (
	"testing"
)

//TODO: table tests
func TestFindPivot(t *testing.T) {
	t.Run("not rotated", func(t *testing.T) {
		arr := []int{0,1,2,4,5,6,7}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 0 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
	t.Run("not rotated", func(t *testing.T) {
		arr := []int{1,2}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 0 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
	t.Run("rotated", func(t *testing.T) {
		arr := []int{2,1}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 1 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
	t.Run("not rotated", func(t *testing.T) {
		arr := []int{1}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 0 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
	t.Run("rotated", func(t *testing.T) {
		arr := []int{3,1,2}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 1 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
	t.Run("rotated", func(t *testing.T) {
		arr := []int{4,5,6,7,0,1,2}
		pivot := FindPivot(arr, 0, len(arr)-1)
		if pivot != 4 {
			t.Fatalf("unexpected pivot %d", pivot)
		}
	})
}

func TestSearch(t *testing.T) {
	t.Run("miss", func(t *testing.T) {
		arr := []int{4,5,6,7,0,1,2}
		res := Search(arr, 3)
		if res != -1 {
			t.Fatalf("unexpected result %d", res)
		}
	})
	t.Run("find", func(t *testing.T) {
		arr := []int{4,5,6,7,0,1,2}
		res := Search(arr, 0)
		if res != 4 {
			t.Fatalf("unexpected result %d", res)
		}
	})
	t.Run("find", func(t *testing.T) {
		arr := []int{1,3}
		res := Search(arr, 3)
		if res != 1 {
			t.Fatalf("unexpected result %d", res)
		}
	})
	t.Run("miss", func(t *testing.T) {
		arr := []int{3, 1}
		res := Search(arr, 3)
		if res != 0 {
			t.Fatalf("unexpected result %d", res)
		}
	})
}
