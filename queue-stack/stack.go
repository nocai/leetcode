type stack struct {
	data []int
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) pop() bool {
	if s.isEmpty() {
		return false
	}

	s.data = s.data[:len(s.data)-1]
	return true
}
