package main

func main() {

}

const stackNil = -1

type stack struct {
	data []int
	maxN int
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(val int) {
	if len(s.data) == 0 {
		s.data = append(s.data, val)
		s.maxN = val
	} else {
		if val > s.maxN {
			s.data = append(s.data, 2*val-s.maxN)
			s.maxN = val
		} else {
			s.data = append(s.data, val)
		}
	}
}

func (s *stack) pop() int {
	l := len(s.data)

	if l == 0 {
		return stackNil
	}

	k := s.data[l-1]
	s.data = s.data[:l-1]
	if k > s.maxN {
		k, s.maxN = s.maxN, 2*s.maxN-k
	}

	return k
}

func (s *stack) max() int {
	l := len(s.data)
	if l == 0 {
		return stackNil
	}

	return s.maxN
}
