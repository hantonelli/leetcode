package ds

import (
	"container/list"
)

type Stack struct {
	items *list.List
	size  int
}

func NewStack() *Stack {
	return &Stack{items: list.New()}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(x interface{}) {
	s.items.PushBack(x)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.items.Len() == 0 {
		return nil
	}
	tail := s.items.Back()
	val := tail.Value
	s.items.Remove(tail)
	s.size--
	return val
}

func (s *Stack) Peek() interface{} {
	if s.items.Len() == 0 {
		return nil
	}
	tail := s.items.Back()
	val := tail.Value
	return val
}
