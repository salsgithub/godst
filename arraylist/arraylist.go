package arraylist

import (
	"fmt"
	"slices"
)

const (
	defaultGrowthFactor = 2.0
	defaultShrinkFactor = 0.25
)

type Option[T comparable] func(*List[T])

type List[T comparable] struct {
	elements        []T
	initialCapacity int
	growthFactor    float64
	shrinkFactor    float64
	fixedSize       *int
}

func New[T comparable](options ...Option[T]) *List[T] {
	l := &List[T]{
		growthFactor: defaultGrowthFactor,
		shrinkFactor: defaultShrinkFactor,
		fixedSize:    nil,
	}
	for _, option := range options {
		option(l)
	}
	l.initialCapacity = cap(l.elements)
	if l.elements == nil {
		l.elements = make([]T, 0, l.initialCapacity)
	}
	return l
}

func WithInitialValues[T comparable](values ...T) Option[T] {
	return func(l *List[T]) {
		if len(values) > 0 {
			l.Add(values...)
		}
	}
}

func WithInitialCapacity[T comparable](capacity int) Option[T] {
	return func(l *List[T]) {
		if l.Len() == 0 {
			l.elements = make([]T, l.Len(), capacity)
		} else {
			elements := make([]T, l.Len(), capacity)
			copy(elements, l.elements)
			l.elements = elements
		}
	}
}

func WithGrowthFactor[T comparable](factor float64) Option[T] {
	return func(l *List[T]) {
		if factor > 1.0 {
			l.growthFactor = factor
		}
	}
}

func WithShrinkFactor[T comparable](factor float64) Option[T] {
	return func(l *List[T]) {
		if factor >= 0.0 {
			l.shrinkFactor = factor
		}
	}
}

func WithFixedSize[T comparable](size int) Option[T] {
	return func(l *List[T]) {
		if size >= 0 {
			l.fixedSize = &size
		}
	}
}

func (l *List[T]) Add(values ...T) {
	addition := len(values)
	if addition <= 0 {
		return
	}
	if l.IsFixedSize() {
		if l.Len() >= *l.fixedSize {
			return
		}
		remainingSize := *l.fixedSize - l.Len()
		if addition > remainingSize {
			addition = remainingSize
		}
	}
	newSize := l.Len() + addition
	if newSize > cap(l.elements) {
		growth := float64(cap(l.elements)) * l.growthFactor
		newCapacity := int(max(float64(newSize), growth))
		if l.fixedSize != nil && newCapacity > *l.fixedSize {
			newCapacity = *l.fixedSize
		}
		l.grow(newCapacity)
	}
	l.elements = append(l.elements, values[:addition]...)
}

func (l *List[T]) checkBounds(index int) error {
	if index < 0 || index >= l.Len() {
		return fmt.Errorf("%w: %d", ErrIndexOutOfBounds, index)
	}
	return nil
}

func (l *List[T]) Remove(index int) error {
	if err := l.checkBounds(index); err != nil {
		return err
	}
	l.elements = slices.Delete(l.elements, index, index+1)
	l.shrink()
	return nil
}

func (l *List[T]) Replace(index int, value T) error {
	if err := l.checkBounds(index); err != nil {
		return err
	}
	l.elements[index] = value
	return nil
}

func (l *List[T]) Contains(value T) bool {
	return slices.Contains(l.elements, value)
}

func (l *List[T]) ContainsAll(values ...T) bool {
	if len(values) == 0 {
		return true
	}
	for _, v := range values {
		if !l.Contains(v) {
			return false
		}
	}
	return true
}

func (l *List[T]) Sort(less func(a, b T) int) {
	slices.SortFunc(l.elements, less)
}

func (l *List[T]) Clear() {
	clear(l.elements)
	l.elements = l.elements[:0]
}

func (l *List[T]) Reset() {
	l.elements = make([]T, 0, l.initialCapacity)
}

func (l *List[T]) Reverse() {
	slices.Reverse(l.elements)
}

func (l *List[T]) Value(index int) (T, error) {
	if err := l.checkBounds(index); err != nil {
		var zero T
		return zero, err
	}
	return l.elements[index], nil
}

func (l *List[T]) Values() []T {
	return slices.Clone(l.elements)
}

func (l *List[T]) IsFixedSize() bool {
	return l.fixedSize != nil
}

func (l *List[T]) Len() int {
	return len(l.elements)
}

func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) grow(capacity int) {
	if cap(l.elements) >= capacity {
		return
	}
	elements := make([]T, l.Len(), capacity)
	copy(elements, l.elements)
	l.elements = elements
}

func (l *List[T]) shrink() {
	if l.shrinkFactor <= 0.0 {
		return
	}
	capacity := cap(l.elements)
	if capacity == 0 {
		return
	}
	if float64(l.Len())/float64(capacity) < l.shrinkFactor {
		l.shrinkToFit()
	}
}

func (l *List[T]) shrinkToFit() {
	elements := make([]T, l.Len())
	copy(elements, l.elements)
	l.elements = elements
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.elements)
}
