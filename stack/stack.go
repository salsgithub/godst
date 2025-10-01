package stack

import (
	"fmt"
	"slices"
)

type Option[T comparable] func(*Stack[T])

type Stack[T comparable] struct {
	contents []T
}

func New[T comparable](options ...Option[T]) *Stack[T] {
	s := &Stack[T]{}
	for _, option := range options {
		option(s)
	}
	if s.contents == nil {
		s.contents = make([]T, 0)
	}
	return s
}

func WithInitialValues[T comparable](values ...T) Option[T] {
	return func(s *Stack[T]) {
		for _, value := range values {
			s.Push(value)
		}
	}
}

func WithInitialCapacity[T comparable](capacity int) Option[T] {
	return func(s *Stack[T]) {
		if s.Size() == 0 {
			s.contents = make([]T, 0, capacity)
		} else {
			contents := make([]T, s.Size(), capacity)
			copy(contents, s.contents)
			s.contents = contents
		}
	}
}

func (s *Stack[T]) Push(values ...T) {
	s.contents = append(s.contents, values...)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrStackEmpty
	}
	index := len(s.contents) - 1
	value := s.contents[index]
	s.contents = s.contents[:index]
	return value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrStackEmpty
	}
	return s.contents[len(s.contents)-1], nil
}

func (s *Stack[T]) Clear() {
	clear(s.contents)
	s.contents = s.contents[:0]
}

func (s *Stack[T]) Reverse() {
	slices.Reverse(s.contents)
}

func (s *Stack[T]) Size() int {
	return len(s.contents)
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.contents)
}

func (l *Stack[T]) IsEmpty() bool {
	return l.Size() == 0
}
