package set

import (
	"fmt"
	"slices"
	"strings"
)

type Set[T comparable] struct {
	contents map[T]struct{}
}

func New[T comparable](values ...T) *Set[T] {
	s := &Set[T]{
		contents: make(map[T]struct{}),
	}
	s.Add(values...)
	return s
}

func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		s.contents[value] = struct{}{}
	}
}

func (s *Set[T]) Remove(value T) {
	delete(s.contents, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.contents[value]
	return ok
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()
	for item := range s.contents {
		union.Add(item)
	}
	for item := range other.contents {
		union.Add(item)
	}
	return union
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := New[T]()
	if s.Len() < other.Len() {
		for item := range s.contents {
			if other.Contains(item) {
				intersection.Add(item)
			}
		}
	} else {
		for item := range other.contents {
			if s.Contains(item) {
				intersection.Add(item)
			}
		}
	}
	return intersection
}

func (s *Set[T]) Delta(other *Set[T]) *Set[T] {
	delta := New[T]()
	for item := range s.contents {
		if !other.Contains(item) {
			delta.Add(item)
		}
	}
	return delta
}

func (s *Set[T]) All(yield func(T) bool) {
	for item := range s.contents {
		if !yield(item) {
			return
		}
	}
}

func (s *Set[T]) Len() int {
	return len(s.contents)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.contents) == 0
}

func (s *Set[T]) Clear() {
	clear(s.contents)
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, s.Len())
	for value := range s.contents {
		values = append(values, value)
	}
	return values
}

func (s *Set[T]) String() string {
	values := s.Values()
	a := any(values)
	switch t := a.(type) {
	case []string:
		slices.Sort(t)
	case []int:
		slices.Sort(t)
	case []float64:
		slices.Sort(t)
	}
	stringValues := make([]string, 0, len(values))
	for _, value := range values {
		stringValues = append(stringValues, fmt.Sprintf("%v", value))
	}
	return strings.Join(stringValues, ", ")
}
