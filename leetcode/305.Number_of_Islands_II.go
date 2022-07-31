package main

var dirs = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func numIslands2(m int, n int, positions [][]int) []int {
	res := make([]int, 0, len(positions))
	var islandsCount int
	uf := NewUnionFind(m * n)

	for _, pos := range positions {
		root := n*pos[0] + pos[1]
		if uf.roots[root] != -1 {
			res = append(res, islandsCount)
			continue
		}
		islandsCount++
		uf.roots[root] = root
		for _, dir := range dirs {
			x := pos[0] + dir[0]
			y := pos[1] + dir[1]
			nb := n*x + y
			if x < 0 || x >= m || y < 0 || y >= n || uf.roots[nb] == -1 {
				continue
			}
			nbRoot := uf.find(nb)

			if root != nbRoot {
				if uf.union(root, nbRoot) {
					islandsCount--
				}
			}
		}
		res = append(res, islandsCount)
	}
	return res
}

type UnionFind struct {
	roots []int
	rank  []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		roots: make([]int, n),
		rank:  make([]int, n),
	}
	for i := range uf.roots {
		uf.roots[i] = -1
	}
	return uf
}

func (uf *UnionFind) union(p, q int) bool {
	pR := uf.find(p)
	qR := uf.find(q)
	if pR == qR {
		return false
	}
	if uf.rank[pR] < uf.rank[qR] {
		uf.roots[pR] = qR
	} else if uf.rank[pR] > uf.rank[qR] {
		uf.roots[qR] = pR
	} else {
		uf.roots[qR] = pR
		uf.rank[pR]++
	}
	return true
}

func (uf *UnionFind) find(x int) int {
	for uf.roots[x] != x {
		uf.roots[x] = uf.roots[uf.roots[x]]
		x = uf.roots[x]
	}
	return x
}
