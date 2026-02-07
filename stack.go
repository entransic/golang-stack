package stack

import "errors"

type Stack struct {
	size   int
	values []string
}

func NewStack() *Stack {
	return &Stack{0, make([]string, 5)}
}

func (s *Stack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(elem string) error {
	ss := s.Size()
	if ss == 4 {
		return errors.New("Stack full")
	}
	s.size++
	s.values[s.size] = elem
	return nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Empty() {
	s.size = 0
	s.values = s.values[:0]
}

func (s *Stack) Pop() (string, error) {
	if s.Size() == 0 {
		return "", errors.New("Poping empty stack")
	}
	popped := s.values[int(s.size)]
	s.values = s.values[:len(s.values)-1]
	s.size--
	return popped, nil
}
