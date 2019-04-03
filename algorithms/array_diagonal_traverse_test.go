package algorithms

import "testing"

var matrix = [][]int{
	{1,2,3},
	{4,5,6},
	{7,8,9},
}

func BenchmarkDiagonalTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDiagonalOrder(matrix)
	}
}
