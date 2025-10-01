package doublylinkedlist

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	value    T
	previous *node[T]
	next     *node[T]
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
		l.head.previous = n
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
		n.previous = l.tail
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
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.previous = nil
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
	} else {
		l.tail = l.tail.previous
		l.tail.next = nil
	}
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
	if index < l.size/2 {
		target := l.head
		for i := 0; i < index-1; i++ {
			target = target.next
		}
		next := target.next
		node := &node[T]{
			value:    value,
			previous: target,
			next:     next,
		}
		target.next = node
		next.previous = node
	} else {
		target := l.tail
		for i := 0; i < l.size-index-1; i++ {
			target = target.previous
		}
		previous := target.previous
		node := &node[T]{
			value:    value,
			previous: previous,
			next:     target,
		}
		target.previous = node
		previous.next = node
	}
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
	if index == l.size-1 {
		return l.PopBack()
	}
	var target *node[T]
	if index < l.size/2 {
		target = l.head
		for range index {
			target = target.next
		}
	} else {
		target = l.tail
		for i := 0; i < l.size-index-1; i++ {
			target = target.previous
		}
	}
	value := target.value
	previous := target.previous
	next := target.next
	previous.next = next
	next.previous = previous
	l.size--
	return value, true
}

func (l *List[T]) At(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}
	var target *node[T]
	if index < l.size/2 {
		target = l.head
		for range index {
			target = target.next
		}
	} else {
		target = l.tail
		for i := 0; i < l.size-index-1; i++ {
			target = target.previous
		}
	}
	return target.value, true
}

func (l *List[T]) Replace(index int, value T) {
	if index < 0 || index >= l.size {
		return
	}
	var target *node[T]
	if index < l.size/2 {
		target = l.head
		for range index {
			target = target.next
		}
	} else {
		target = l.tail
		for i := 0; i < l.size-index-1; i++ {
			target = target.previous
		}
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

func (l *List[T]) ReverseAll(yield func(T) bool) {
	node := l.tail
	for node != nil {
		if !yield(node.value) {
			return
		}
		node = node.previous
	}
}

func (l *List[T]) Reverse() {
	if l.size < 2 {
		return
	}
	node := l.head
	for node != nil {
		next := node.next
		node.next = node.previous
		node.previous = next
		node = next
	}
	l.head, l.tail = l.tail, l.head
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

func (l *List[T]) ReverseValues() []T {
	values := make([]T, 0, l.size)
	node := l.tail
	for node != nil {
		values = append(values, node.value)
		node = node.previous
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
			builder.WriteString(" <-> ")
		}
		node = node.next
	}
	builder.WriteString("]")
	return builder.String()
}
