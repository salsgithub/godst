package stack

import (
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestStack_New(t *testing.T) {
	tests := []struct {
		name             string
		options          []Option[int]
		expectedSize     int
		expectedCapacity int
		expectedContents []int
	}{
		{
			name:             "Empty stack",
			options:          nil,
			expectedSize:     0,
			expectedCapacity: 0,
			expectedContents: []int{},
		},
		{
			name:             "Stack with initial values",
			options:          []Option[int]{WithInitialValues(1, 2, 3, 4)},
			expectedSize:     4,
			expectedCapacity: 4,
			expectedContents: []int{1, 2, 3, 4},
		},
		{
			name:             "Stack with initial capacity",
			options:          []Option[int]{WithInitialCapacity[int](20)},
			expectedSize:     0,
			expectedCapacity: 20,
			expectedContents: []int{},
		},
		{
			name:             "Stack with initial capacity and values",
			options:          []Option[int]{WithInitialValues(1, 2, 3, 4), WithInitialCapacity[int](20)},
			expectedSize:     4,
			expectedCapacity: 20,
			expectedContents: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := New(test.options...)
			size := s.Size()
			capacity := cap(s.contents)
			assertx.Equal(t, size, test.expectedSize)
			assertx.Equal(t, capacity, test.expectedCapacity)
			assertx.Equal(t, s.contents, test.expectedContents)
		})
	}
}

func TestStack_Push(t *testing.T) {
	s := New[int]()
	s.Push(20, 30, 40, 50)
	assertx.Equal(t, s.Size(), 4)
}

func TestStack_Pop(t *testing.T) {
	t.Run("pop from stack yields value", func(t *testing.T) {
		s := New[int]()
		s.Push(20)
		value, err := s.Pop()
		assertx.Nil(t, err)
		assertx.Equal(t, value, 20)
	})
	t.Run("pop from empty stack yields error", func(t *testing.T) {
		s := New[int]()
		value, err := s.Pop()
		assertx.NotNil(t, err)
		assertx.ErrorIs(t, err, ErrStackEmpty)
		assertx.Equal(t, value, 0)
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("peek yields last pushed value", func(t *testing.T) {
		s := New[string]()
		s.Push("hello")
		s.Push("go")
		value, err := s.Peek()
		assertx.Nil(t, err)
		assertx.Equal(t, value, "go")
	})
	t.Run("peek empty stack yields error", func(t *testing.T) {
		s := New[string]()
		value, err := s.Peek()
		assertx.NotNil(t, err)
		assertx.ErrorIs(t, err, ErrStackEmpty)
		assertx.Equal(t, value, "")
	})
}

func TestStack_Reverse(t *testing.T) {
	s := New(WithInitialValues("language", "go"))
	s.Reverse()
	assertx.Equal(t, s.contents, []string{"go", "language"})
}

func TestStack_Clear(t *testing.T) {
	s := New(WithInitialValues("a", "b", "c"))
	initialCapacity := cap(s.contents)
	s.Clear()
	assertx.Equal(t, s.Size(), 0)
	assertx.Equal(t, cap(s.contents), initialCapacity)
	assertx.True(t, s.IsEmpty())
}

func TestStack_String(t *testing.T) {
	s := New(WithInitialValues("a", "b", "c"))
	value := s.String()
	assertx.Equal(t, value, "[a b c]")
}
