package ds

// nolint

// https://github.com/zyedidia/generic/blob/master/stack/stack.go

// Stack3 implements a LIFO stack with peeking.
type Stack3[T any] struct {
	entries []T
}

// New returns an empty stack.
func NewStack3[T any]() *Stack3[T] {
	return &Stack3[T]{
		entries: nil,
	}
}

// Push places 'value' at the top of the stack.
func (s *Stack3[T]) Push(value T) {
	s.entries = append(s.entries, value)
}

// Pop removes the stack's top element and returns it. If the stack is empty it
// returns the zero value.
func (s *Stack3[T]) Pop() (t T) {
	if len(s.entries) == 0 {
		return t
	}
	v := s.entries[len(s.entries)-1]
	s.entries = s.entries[:len(s.entries)-1]
	return v
}

// Peek returns the stack's top element but does not remove it. If the stack is
// empty the zero value is returned.
func (s *Stack3[T]) Peek() (t T) {
	if len(s.entries) == 0 {
		return t
	}
	return s.entries[len(s.entries)-1]
}

// Size returns the number of elements in the stack.
func (s *Stack3[T]) Size() int {
	return len(s.entries)
}

// Copy returns a copy of this stack.
func (s *Stack3[T]) Copy() *Stack3[T] {
	entries := make([]T, len(s.entries))
	copy(entries, s.entries)
	return &Stack3[T]{
		entries: entries,
	}
}

// IsEmpty returns true if the stack is empty
func (s *Stack3[T]) IsEmpty() bool {
	return len(s.entries) == 0
}

// Clear empty all entries
func (s *Stack3[T]) Clear() {
	s.entries = nil
}
