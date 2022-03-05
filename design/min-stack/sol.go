package minstack

import "math"

type MinStack struct {
	entries []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(value int) {
	s.entries = append(s.entries, value)
}

func (s *MinStack) Pop() {
	if len(s.entries) == 0 {
		return
	}
	s.entries = s.entries[:len(s.entries)-1]
}

func (s *MinStack) Top() int {
	if len(s.entries) == 0 {
		return -1
	}
	return s.entries[len(s.entries)-1]

}

func (s *MinStack) GetMin() int {
	min := math.MaxInt32
	for i := 0; i < len(s.entries); i++ {
		if s.entries[i] < min {
			min = s.entries[i]
		}
	}
	return min
}
