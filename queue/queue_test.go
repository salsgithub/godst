package queue

import (
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestQueue_Enqueue(t *testing.T) {
	q := New[string]()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")
	assertx.Equal(t, q.Len(), 3)
	assertx.False(t, q.IsEmpty())
}

func TestQueue_Dequeue(t *testing.T) {
	q := New[string]()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")
	values := make([]string, 0, 3)
	for !q.IsEmpty() {
		value, ok := q.Dequeue()
		if !ok {
			t.Error("failed to dequeue")
		}
		values = append(values, value)
	}
	assertx.Equal(t, values, []string{"A", "B", "C"})
	assertx.Equal(t, q.Len(), 0)
}

func TestQueue_Peek(t *testing.T) {
	q := New[string]()
	q.Enqueue("Go")
	peek, ok := q.Peek()
	assertx.True(t, ok)
	assertx.Equal(t, peek, "Go")
	assertx.Equal(t, q.Len(), 1)
}

func TestQueue_Clear(t *testing.T) {
	q := New[string]()
	q.Enqueue("Go")
	q.Clear()
	assertx.Equal(t, q.Len(), 0)
	assertx.True(t, q.IsEmpty())
}
