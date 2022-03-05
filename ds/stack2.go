package ds

type element struct {
	data interface{}
	next *element
}

type Stack2 struct {
	head *element
	size int
}

func NewStack2() *Stack2 {
	return new(Stack2)
}

func (s *Stack2) Size() int {
	return s.size
}

func (s *Stack2) Push(data interface{}) {
	element := new(element)
	element.data = data
	temp := s.head
	element.next = temp
	s.head = element
	s.size++

}

func (s *Stack2) Pop() interface{} {
	if s.head == nil {
		return nil
	}
	r := s.head.data
	s.head = s.head.next
	s.size--

	return r
}
