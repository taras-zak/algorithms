package main

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	res := 1
	for len(s.stack) > 0 {
		top := s.stack[len(s.stack)-1]
		if top[0] <= price {
			s.stack = s.stack[:len(s.stack)-1]
			res += top[1]
		} else {
			break
		}
	}
	s.stack = append(s.stack, [2]int{price, res})
	return res
}
