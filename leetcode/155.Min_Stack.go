package main

import (
	"fmt"
)

type MinStack struct {
	values    []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{values: make([]int, 0), minValues: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.values = append(s.values, val)
	if len(s.minValues) == 0 {
		s.minValues = append(s.minValues, val)
		return
	}

	minValue := val
	if s.minValues[len(s.minValues)-1] < minValue {
		minValue = s.minValues[len(s.minValues)-1]
	}
	s.minValues = append(s.minValues, minValue)
	return
}

func (s *MinStack) Pop() {
	s.values = s.values[:len(s.values)-1]
	s.minValues = s.minValues[:len(s.minValues)-1]
}

func (s *MinStack) Top() int {
	return s.values[len(s.values)-1]
}

func (s *MinStack) GetMin() int {
	return s.minValues[len(s.minValues)-1]
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())

	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}
