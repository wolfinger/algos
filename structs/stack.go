package stack

type Stack struct {
	stack []int
}

func (s *Stack) Push(item int) {
	// add to stack
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() int {
	// remove from stack
	item := s.stack[len(s.stack)-1]
	s.stack = append(s.stack[:len(s.stack)-1])

	return item
}

func (s *Stack) Read() (int, bool) {
	// read from stack
	if len(s.stack) == 0 {
		return 0, true
	} else {
		return s.stack[len(s.stack)-1], false
	}
}
