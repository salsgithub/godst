package queue

import "github.com/salsgithub/godst/doublylinkedlist"

type Queue[T comparable] struct {
	list *doublylinkedlist.List[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		list: doublylinkedlist.New[T](),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.list.PopFront()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Front()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}
