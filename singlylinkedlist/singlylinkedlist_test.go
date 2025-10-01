package singlylinkedlist

import (
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestSinglyLinkedList_New(t *testing.T) {
	l := New(1, 2, 3, 4, 5)
	head, ok := l.Front()
	assertx.True(t, ok)
	assertx.Equal(t, head, 1)
	back, ok := l.Back()
	assertx.True(t, ok)
	assertx.Equal(t, back, 5)
	assertx.Equal(t, l.size, 5)
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	l := New[int]()
	l.Prepend(1)
	l.Prepend(2)
	assertx.Equal(t, l.Len(), 2)
	expectedHead := &node[int]{
		value: 2,
		next: &node[int]{
			value: 1,
			next:  nil,
		},
	}
	expectedTail := &node[int]{
		value: 1,
		next:  nil,
	}
	assertx.Equal(t, l.head, expectedHead)
	assertx.Equal(t, l.tail, expectedTail)
}

func TestSinglyLinkedList_Append(t *testing.T) {
	l := New[int]()
	l.Append(1)
	l.Append(2)
	assertx.Equal(t, l.Len(), 2)
	expectedHead := &node[int]{
		value: 1,
		next: &node[int]{
			value: 2,
			next:  nil,
		},
	}
	expectedTail := &node[int]{
		value: 2,
		next:  nil,
	}
	assertx.Equal(t, l.head, expectedHead)
	assertx.Equal(t, l.tail, expectedTail)
}

func TestSinglyLinkedList_Front(t *testing.T) {
	t.Run("front with empty list", func(t *testing.T) {
		l := New[string]()
		front, ok := l.Front()
		assertx.False(t, ok)
		assertx.Equal(t, front, "")
	})
	t.Run("front with multiple nodes", func(t *testing.T) {
		l := New[string]()
		l.Prepend("Go")
		l.Prepend("language")
		front, ok := l.Front()
		assertx.True(t, ok)
		assertx.Equal(t, front, "language")
	})
}

func TestSinglyLinkedList_PopFront(t *testing.T) {
	t.Run("pop front on empty list", func(t *testing.T) {
		l := New[int]()
		front, ok := l.PopFront()
		assertx.False(t, ok)
		assertx.Equal(t, front, 0)
		assertx.True(t, l.IsEmpty())
	})
	t.Run("pop front with one node in list", func(t *testing.T) {
		l := New[int]()
		l.Prepend(1)
		front, ok := l.PopFront()
		assertx.True(t, ok)
		assertx.Equal(t, front, 1)
		assertx.Nil(t, l.head)
		assertx.Nil(t, l.tail)
		assertx.Equal(t, l.size, 0)
	})
	t.Run("pop front with multiple nodes in list", func(t *testing.T) {
		l := New[int]()
		l.Prepend(1)
		l.Prepend(2)
		front, ok := l.PopFront()
		assertx.True(t, ok)
		assertx.Equal(t, front, 2)
		assertx.Equal(t, l.size, 1)
	})
}

func TestSinglyLinkedList_Back(t *testing.T) {
	t.Run("back with empty list", func(t *testing.T) {
		l := New[int]()
		back, ok := l.Back()
		assertx.False(t, ok)
		assertx.Equal(t, back, 0)
	})
	t.Run("back with one node", func(t *testing.T) {
		l := New[string]()
		l.Prepend("Go")
		back, ok := l.Back()
		assertx.True(t, ok)
		assertx.Equal(t, back, "Go")
	})
	t.Run("back with multiple nodes", func(t *testing.T) {
		l := New[string]()
		l.Prepend("Go")
		l.Prepend("language")
		back, ok := l.Back()
		assertx.True(t, ok)
		assertx.Equal(t, back, "Go")
	})
}

func TestSinglyLinkedList_PopBack(t *testing.T) {
	t.Run("pop back on empty list", func(t *testing.T) {
		l := New[string]()
		back, ok := l.PopBack()
		assertx.False(t, ok)
		assertx.Equal(t, back, "")
	})
	t.Run("pop back with one node in list", func(t *testing.T) {
		l := New[string]()
		l.Append("Go")
		back, ok := l.PopBack()
		assertx.True(t, ok)
		assertx.Equal(t, back, "Go")
		assertx.Equal(t, l.size, 0)
	})
	t.Run("pop back with multiple nodes in list", func(t *testing.T) {
		l := New[string]()
		l.Prepend("using")
		l.Prepend("language")
		l.Prepend("Go")
		back, ok := l.PopBack()
		assertx.True(t, ok)
		assertx.Equal(t, back, "using")
		assertx.Equal(t, l.size, 2)
	})
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	t.Run("insert out of bounds", func(t *testing.T) {
		l := New[int]()
		inserted := l.Insert(1, 1)
		assertx.False(t, inserted)
	})
	t.Run("insert into empty list", func(t *testing.T) {
		l := New[int]()
		inserted := l.Insert(0, 0)
		assertx.True(t, inserted)
	})
	t.Run("insert at head", func(t *testing.T) {
		l := New(1, 2, 3)
		inserted := l.Insert(0, 10)
		assertx.True(t, inserted)
		front, ok := l.Front()
		assertx.True(t, ok)
		assertx.Equal(t, front, 10)
	})
	t.Run("insert at tail", func(t *testing.T) {
		l := New(1, 2)
		inserted := l.Insert(2, 20)
		assertx.True(t, inserted)
		back, ok := l.Back()
		assertx.True(t, ok)
		assertx.Equal(t, back, 20)
		assertx.Equal(t, l.Len(), 3)
	})
	t.Run("insert anywhere in the list", func(t *testing.T) {
		l := New(1, 2, 3, 4)
		inserted := l.Insert(2, 21)
		assertx.True(t, inserted)
		assertx.Equal(t, l.Values(), []int{1, 2, 21, 3, 4})
	})
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	tests := []struct {
		name           string
		values         []int
		removalIndex   int
		expectedValue  int
		expectedRemove bool
		expectedValues []int
	}{
		{
			name:           "Remove from empty list does not remove anything",
			values:         []int{},
			removalIndex:   0,
			expectedValue:  0,
			expectedRemove: false,
			expectedValues: []int{},
		},
		{
			name:           "Remove from invalid index does not remove anything",
			values:         []int{1},
			removalIndex:   2,
			expectedValue:  0,
			expectedRemove: false,
			expectedValues: []int{1},
		},
		{
			name:           "Remove first node removes the head",
			values:         []int{1},
			removalIndex:   0,
			expectedValue:  1,
			expectedRemove: true,
			expectedValues: []int{},
		},
		{
			name:           "Remove at valid index removes that node",
			values:         []int{1, 2, 3, 4, 5},
			removalIndex:   3,
			expectedValue:  4,
			expectedRemove: true,
			expectedValues: []int{1, 2, 3, 5},
		},
		{
			name:           "Remove the last node removes that node",
			values:         []int{1, 2, 3, 4, 5},
			removalIndex:   4,
			expectedValue:  5,
			expectedRemove: true,
			expectedValues: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(test.values...)
			value, ok := l.Remove(test.removalIndex)
			assertx.Equal(t, ok, test.expectedRemove)
			assertx.Equal(t, value, test.expectedValue)
			assertx.Equal(t, l.Values(), test.expectedValues)
		})
	}
}

func TestSinglyLinkedList_At(t *testing.T) {
	t.Run("invalid index provided for at", func(t *testing.T) {
		l := New(1, 2, 3)
		value, ok := l.At(10)
		assertx.False(t, ok)
		assertx.Equal(t, value, 0)
		assertx.Equal(t, l.Len(), 3)
	})
	t.Run("valid index provided for at", func(t *testing.T) {
		l := New("the", "language", "is", "Go")
		value, ok := l.At(3)
		assertx.True(t, ok)
		assertx.Equal(t, value, "Go")
		assertx.Equal(t, l.Len(), 4)
	})
}

func TestSinglyLinkedList_Replace(t *testing.T) {
	t.Run("invalid index provided for replace", func(t *testing.T) {
		l := New(1, 2, 3)
		l.Replace(10, 10)
		assertx.Equal(t, l.Values(), []int{1, 2, 3})
	})
	t.Run("valid index provided for replace", func(t *testing.T) {
		l := New("the", "best", "language", "is", "TypeScript")
		l.Replace(4, "Go")
		value, ok := l.At(4)
		assertx.True(t, ok)
		assertx.Equal(t, value, "Go")
	})
}

func TestSinglyLinkedList_Index(t *testing.T) {
	t.Run("value not found", func(t *testing.T) {
		l := New(1, 2, 3)
		index := l.Index(20)
		assertx.Equal(t, index, -1)
	})
	t.Run("value found", func(t *testing.T) {
		l := New(1, 2, 3)
		index := l.Index(3)
		assertx.Equal(t, index, 2)
	})
}

func TestSinglyLinkedList_All(t *testing.T) {
	t.Run("iterate over empty list", func(t *testing.T) {
		l := New[int]()
		var results []int
		for value := range l.All {
			results = append(results, value)
			t.Error("iterator for an empty list should not yield any values")
		}
		assertx.Nil(t, results)
	})
	t.Run("iterate over a list with one node", func(t *testing.T) {
		l := New(1)
		var results []int
		for value := range l.All {
			results = append(results, value)
		}
		assertx.Equal(t, results, []int{1})
	})
	t.Run("iterate over a list with multiple nodes", func(t *testing.T) {
		l := New("can", "you", "guess", "the", "language")
		var results []string
		for value := range l.All {
			results = append(results, value)
		}
		assertx.Equal(t, results, []string{"can", "you", "guess", "the", "language"})
	})
	t.Run("iterate using all directly", func(t *testing.T) {
		l := New(1, 2, 3)
		var results []int
		l.All(func(i int) bool {
			if i > 1 {
				results = append(results, i)
				return true
			}
			return false
		})
		assertx.Nil(t, results)
		l.All(func(i int) bool {
			if i > 0 {
				results = append(results, i)
			}
			return i > 0
		})
		assertx.Equal(t, results, []int{1, 2, 3})
	})
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	t.Run("reversing a list with size less than two has no effect", func(t *testing.T) {
		l := New(1)
		l.Reverse()
		assertx.Equal(t, l.Values(), []int{1})
	})
	t.Run("reversing a list", func(t *testing.T) {
		l := New("language", "the", "is", "Go")
		l.Reverse()
		assertx.Equal(t, l.Values(), []string{"Go", "is", "the", "language"})
	})
}

func TestSinglyLinkedList_Values(t *testing.T) {
	l := New(1, 2, 3)
	values := l.Values()
	assertx.Equal(t, values, []int{1, 2, 3})
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	l := New(1, 2, 3, 4, 5, 6, 7)
	l.Clear()
	assertx.Equal(t, l.Len(), 0)
	assertx.Nil(t, l.head)
	assertx.Nil(t, l.tail)
}

func TestSinglyLinkedList_String(t *testing.T) {
	t.Run("string with empty list", func(t *testing.T) {
		l := New[int]()
		s := l.String()
		assertx.Equal(t, s, "[]")
	})
	t.Run("string with one node in list", func(t *testing.T) {
		l := New[int]()
		l.Append(1)
		s := l.String()
		assertx.Equal(t, s, "[1]")
	})
	t.Run("string with multiple nodes in list", func(t *testing.T) {
		l := New(1, 2, 3, 4, 5, 6, 7)
		s := l.String()
		assertx.Equal(t, s, "[1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7]")
	})
}

func BenchmarkAppend_100(b *testing.B) {
	size := 100
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Append(i)
		}
	}
}

func BenchmarkAppend_1_000(b *testing.B) {
	size := 1_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Append(i)
		}
	}
}

func BenchmarkAppend_10_000(b *testing.B) {
	size := 10_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Append(i)
		}
	}
}

func BenchmarkAppend_100_000(b *testing.B) {
	size := 100_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Append(i)
		}
	}
}

func BenchmarkPrepend_100(b *testing.B) {
	size := 100
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Prepend(i)
		}
	}
}

func BenchmarkPrepend_1_000(b *testing.B) {
	size := 1_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Prepend(i)
		}
	}
}

func BenchmarkPrepend_10_000(b *testing.B) {
	size := 10_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Prepend(i)
		}
	}
}

func BenchmarkPrepend_100_000(b *testing.B) {
	size := 100_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Prepend(i)
		}
	}
}
