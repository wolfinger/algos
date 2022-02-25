package ds

type Queue struct {
	queue []int
}

func (s *Queue) Enqueue(item int) {
	// add to queue
	s.queue = append(s.queue, item)
}

func (s *Queue) Dequeue() (int, bool) {
	// remove from queue
	if len(s.queue) > 0 {
		item := s.queue[0]
		s.queue = s.queue[1:len(s.queue)]

		return item, false
	} else {
		return 0, true
	}
}

func (s *Queue) Read() (int, bool) {
	// read from queue
	if len(s.queue) == 0 {
		return 0, true
	} else {
		return s.queue[0], false
	}
}
