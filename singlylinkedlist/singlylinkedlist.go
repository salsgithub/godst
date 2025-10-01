package singlylinkedlist

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

type List[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T comparable](values ...T) *List[T] {
	l := &List[T]{}
	for _, value := range values {
		l.Append(value)
	}
	return l
}

func (l *List[T]) Prepend(value T) {
	n := &node[T]{value: value}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.size++
}

func (l *List[T]) Append(value T) {
	n := &node[T]{value: value}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func (l *List[T]) Front() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.value, true
}

func (l *List[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	value := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return value, true
}

func (l *List[T]) Back() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.value, true
}

func (l *List[T]) PopBack() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	value := l.tail.value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}
	current := l.head
	for current.next != l.tail {
		current = current.next
	}
	current.next = nil
	l.tail = current
	l.size--
	return value, true
}

func (l *List[T]) Insert(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}
	if index == 0 {
		l.Prepend(value)
		return true
	}
	if index == l.size {
		l.Append(value)
		return true
	}
	target := l.head
	for i := 0; i < index-1; i++ {
		target = target.next
	}
	node := &node[T]{
		value: value,
		next:  target.next,
	}
	target.next = node
	l.size++
	return true
}

func (l *List[T]) Remove(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}
	if index == 0 {
		return l.PopFront()
	}
	target := l.head
	for i := 0; i < index-1; i++ {
		target = target.next
	}
	removal := target.next
	value := removal.value
	target.next = removal.next
	if target.next == nil {
		l.tail = target
	}
	l.size--
	return value, true
}

func (l *List[T]) At(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}
	target := l.head
	for range index {
		target = target.next
	}
	return target.value, true
}

func (l *List[T]) Replace(index int, value T) {
	if index < 0 || index >= l.size {
		return
	}
	target := l.head
	for range index {
		target = target.next
	}
	target.value = value
}

func (l *List[T]) Index(value T) int {
	target := l.head
	index := 0
	for target != nil {
		if target.value == value {
			return index
		}
		target = target.next
		index++
	}
	return -1
}

func (l *List[T]) All(yield func(T) bool) {
	node := l.head
	for node != nil {
		if !yield(node.value) {
			return
		}
		node = node.next
	}
}

func (l *List[T]) Reverse() {
	if l.size < 2 {
		return
	}
	var previous *node[T] = nil
	current := l.head
	l.tail = current
	for current != nil {
		next := current.next
		current.next = previous
		previous = current
		current = next
	}
	l.head = previous
}

func (l *List[T]) Values() []T {
	values := make([]T, 0, l.size)
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	builder := strings.Builder{}
	builder.WriteString("[")
	node := l.head
	for node != nil {
		builder.WriteString(fmt.Sprintf("%v", node.value))
		if node.next != nil {
			builder.WriteString(" -> ")
		}
		node = node.next
	}
	builder.WriteString("]")
	return builder.String()
}
