package searching_test

import (
	"github.com/taras-zak/algorithms/algorithms/searching"
	"testing"
)

// TODO: benchmarks
func TestBinarySearch(t *testing.T) {
	t.Run("find existing odd", func(t *testing.T) {
		arr := []int{0,1,2,4,5,6,7}
		indx := searching.BinarySearch(arr, 0, len(arr)-1, 6)
		if indx != 5 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find existing even", func(t *testing.T) {
		arr := []int{1,2,4,5,6,7}
		indx := searching.BinarySearch(arr, 0, len(arr)-1, 2)
		if indx != 1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find unexisting", func(t *testing.T) {
		arr := []int{0,1,2,4,5,6,7}
		indx := searching.BinarySearch(arr, 0, len(arr)-1, 3)
		if indx != -1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find existing", func(t *testing.T) {
		arr := []int{-1,0,3,5,9,12}
		indx := searching.BinarySearch(arr, 0, len(arr)-1, 13)
		if indx != -1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
}

func TestBinarySearchIter(t *testing.T) {
	t.Run("find existing odd", func(t *testing.T) {
		arr := []int{0,1,2,4,5,6,7}
		indx := searching.BinarySearchIter(arr,  6)
		if indx != 5 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find existing even", func(t *testing.T) {
		arr := []int{1,2,4,5,6,7}
		indx := searching.BinarySearchIter(arr, 2)
		if indx != 1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find unexisting", func(t *testing.T) {
		arr := []int{0,1,2,4,5,6,7}
		indx := searching.BinarySearchIter(arr, 3)
		if indx != -1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
	t.Run("find existing", func(t *testing.T) {
		arr := []int{-1,0,3,5,9,12}
		indx := searching.BinarySearchIter(arr, 13)
		if indx != -1 {
			t.Fatalf("unexpected indx %d", indx)
		}
	})
}
