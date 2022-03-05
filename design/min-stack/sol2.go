package minstack

type Node struct {
	val int
	min int
}
type MinStack2 struct {
	entries []Node
}

func Constructor2() MinStack2 {
	return MinStack2{}
}

func (s *MinStack2) Push(value int) {
	min := value
	if 0 < len(s.entries) && s.entries[len(s.entries)-1].min < value {
		min = s.entries[len(s.entries)-1].min
	}
	s.entries = append(s.entries, Node{val: value, min: min})
}

func (s *MinStack2) Pop() {
	if len(s.entries) == 0 {
		return
	}
	s.entries = s.entries[:len(s.entries)-1]
}

func (s *MinStack2) Top() int {
	if len(s.entries) == 0 {
		return -1
	}
	return s.entries[len(s.entries)-1].val

}

func (s *MinStack2) GetMin() int {
	if len(s.entries) == 0 {
		return -1
	}
	return s.entries[len(s.entries)-1].min
}
