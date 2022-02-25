package ds

type Stack struct {
	stack []int
}

func (s *Stack) Push(item int) {
	// add to stack
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() (int, bool) {
	// remove from stack
	if len(s.stack) > 0 {
		item := s.stack[len(s.stack)-1]
		s.stack = append(s.stack[:len(s.stack)-1])

		return item, false
	} else {
		return 0, true
	}
}

func (s *Stack) Read() (int, bool) {
	// read from stack
	if len(s.stack) == 0 {
		return 0, true
	} else {
		return s.stack[len(s.stack)-1], false
	}
}
