package binarysearchtree

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type node[T Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

type Tree[T Ordered] struct {
	root *node[T]
	size int
}

func New[T Ordered](values ...T) *Tree[T] {
	t := &Tree[T]{}
	for _, value := range values {
		t.Insert(value)
	}
	return t
}

func (t *Tree[T]) Insert(value T) {
	t.root = insert(t.root, value, &t.size)
}

func (t *Tree[T]) InsertAll(values ...T) {
	for _, value := range values {
		t.Insert(value)
	}
}

func insert[T Ordered](n *node[T], value T, size *int) *node[T] {
	if n == nil {
		*size++
		return &node[T]{value: value}
	}
	if value < n.value {
		n.left = insert(n.left, value, size)
	} else if value > n.value {
		n.right = insert(n.right, value, size)
	}
	return n
}

func (t *Tree[T]) Remove(value T) {
	t.root = remove(t.root, value, &t.size)
}

func (t *Tree[T]) RemoveAll(values ...T) {
	for _, value := range values {
		t.Remove(value)
	}
}

func remove[T Ordered](n *node[T], value T, size *int) *node[T] {
	if n == nil {
		return nil
	}
	if value < n.value {
		n.left = remove(n.left, value, size)
	} else if value > n.value {
		n.right = remove(n.right, value, size)
	} else {
		*size--
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		min := minNode(n.right)
		n.value = min.value
		n.right = remove(n.right, min.value, size)
		// extra remove above causes size decrement, bump up the size for correction
		*size++
	}
	return n
}

func minNode[T Ordered](n *node[T]) *node[T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *Tree[T]) Height() int {
	return calculateHeight(t.root)
}

func calculateHeight[T Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	left := calculateHeight(n.left)
	right := calculateHeight(n.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func (t *Tree[T]) Min() (T, bool) {
	if t.root == nil {
		var zero T
		return zero, false
	}
	n := t.root
	for n.left != nil {
		n = n.left
	}
	return n.value, true
}

func (t *Tree[T]) Max() (T, bool) {
	if t.root == nil {
		var zero T
		return zero, false
	}
	n := t.root
	for n.right != nil {
		n = n.right
	}
	return n.value, true
}

func (t *Tree[T]) Contains(value T) bool {
	return contains(t.root, value)
}

func contains[T Ordered](n *node[T], value T) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	}
	if value < n.value {
		return contains(n.left, value)
	}
	return contains(n.right, value)
}

func (t *Tree[T]) ValuesPreOrder() []T {
	var values []T
	traversePreOrder(t.root, &values)
	return values
}

func traversePreOrder[T Ordered](start *node[T], values *[]T) {
	if start == nil {
		return
	}
	*values = append(*values, start.value)
	traversePreOrder(start.left, values)
	traversePreOrder(start.right, values)
}

func (t *Tree[T]) ValuesInOrder() []T {
	var values []T
	traverseInOrder(t.root, &values)
	return values
}

func traverseInOrder[T Ordered](start *node[T], values *[]T) {
	if start == nil {
		return
	}
	traverseInOrder(start.left, values)
	*values = append(*values, start.value)
	traverseInOrder(start.right, values)
}

func (t *Tree[T]) ValuesPostOrder() []T {
	var values []T
	traversePostOrder(t.root, &values)
	return values
}

func traversePostOrder[T Ordered](start *node[T], values *[]T) {
	if start == nil {
		return
	}
	traversePostOrder(start.left, values)
	traversePostOrder(start.right, values)
	*values = append(*values, start.value)
}

func (t *Tree[T]) ValuesBreadthFirst() []T {
	if t.root == nil {
		return []T{}
	}
	var values []T
	queue := []*node[T]{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		values = append(values, node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return values
}

func (t *Tree[T]) Len() int {
	return t.size
}

func (t *Tree[T]) IsEmpty() bool {
	return t.size == 0
}
