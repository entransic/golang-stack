package stack

import "errors"

type Stack[T any] struct {
	size   int
	values []T
}

func (s *Stack[T]) NewStack() *Stack[T] {
	return &Stack[T]{0, make([]T, 5)}
}

func (s *Stack[T]) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Push(elem T) error {
	ss := s.Size()
	if ss == 4 {
		return errors.New("Stack full")
	}
	s.size++
	s.values[s.size] = elem
	return nil
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Empty() {
	s.size = 0
	s.values = s.values[:0]
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		return s.NewStack().values[0], errors.New("Popping empty stack")
	}
	popped := s.values[int(s.size)]
	s.values = s.values[:len(s.values)-1]
	s.size--
	return popped, nil
}
