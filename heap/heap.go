package heap

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Heap[T any] struct {
	contents   []T
	comparator func(a, b T) bool
}

func New[T any](comparator func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
	}
}

func NewMin[T Ordered]() *Heap[T] {
	return &Heap[T]{
		comparator: func(a, b T) bool {
			return a < b
		},
	}
}

func NewMax[T Ordered]() *Heap[T] {
	return &Heap[T]{
		comparator: func(a, b T) bool {
			return a > b
		},
	}
}

func (h *Heap[T]) Push(value T) {
	h.contents = append(h.contents, value)
	h.bubbleUp(len(h.contents) - 1)
}

func (h *Heap[T]) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if !h.comparator(h.contents[index], h.contents[parent]) {
			break
		}
		h.contents[index], h.contents[parent] = h.contents[parent], h.contents[index]
		index = parent
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	if h.IsEmpty() {
		var zero T
		return zero, false
	}
	root := h.contents[0]
	last := len(h.contents) - 1
	h.contents[0] = h.contents[last]
	h.contents = h.contents[:last]
	h.bubbleDown(0)
	return root, true
}

func (h *Heap[T]) bubbleDown(index int) {
	for {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		swap := index
		if leftIndex < h.Len() && h.comparator(h.contents[leftIndex], h.contents[swap]) {
			swap = leftIndex
		}
		if rightIndex < h.Len() && h.comparator(h.contents[rightIndex], h.contents[swap]) {
			swap = rightIndex
		}
		if swap == index {
			break
		}
		h.contents[index], h.contents[swap] = h.contents[swap], h.contents[index]
		index = swap
	}
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.IsEmpty() {
		var zero T
		return zero, false
	}
	return h.contents[0], true
}

func (h *Heap[T]) Values() []T {
	values := make([]T, h.Len())
	copy(values, h.contents)
	return values
}

func (h *Heap[T]) Clear() {
	h.contents = h.contents[:0]
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.contents) == 0
}

func (h *Heap[T]) Len() int {
	return len(h.contents)
}
