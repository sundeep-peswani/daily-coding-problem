package main

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

type stream struct {
	data [][]byte
	pos  int
}

func newStream(n int) *stream {
	var s stream
	p := make([]byte, 10)

	for i := 0; i < n; i++ {
		if _, err := rand.Read(p); err == nil {
			s.data = append(s.data, p[0:10])
		} else {
			i--
		}
	}

	return &s
}

func (s *stream) Read(p []byte) (n int, err error) {
	if s.pos >= len(s.data) {
		return 0, errors.New("EOF")
	}

	d := s.data[s.pos]
	for i := 0; i < len(d); i++ {
		p[i] = d[i]
	}

	s.pos++

	return len(d), nil
}

func (s *stream) Has(p []byte) bool {
	for i := 0; i < len(s.data); i++ {
		if equal(p, s.data[i]) {
			return true
		}
	}

	return false
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestPickUniformFromStream(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100000; i *= 10 {
		s := newStream(i)
		actual := pickUniformFromStream(s)

		if !s.Has(actual) {
			t.Errorf("Testing with %d elements: actual %#v not in original data set: %#v\n", i, actual, s.data)
		}
	}
}
