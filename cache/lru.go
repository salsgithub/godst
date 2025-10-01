package cache

type entry[K comparable, V any] struct {
	key   K
	value V
}

type node[K comparable, V any] struct {
	e        *entry[K, V]
	previous *node[K, V]
	next     *node[K, V]
}

type LRU[K comparable, V any] struct {
	cache    map[K]*node[K, V]
	capacity int
	head     *node[K, V]
	tail     *node[K, V]
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		capacity = 1
	}
	head := &node[K, V]{}
	tail := &node[K, V]{}
	head.next = tail
	tail.previous = head
	return &LRU[K, V]{
		cache:    make(map[K]*node[K, V]),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	node, ok := c.cache[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.unlink(node)
	c.addFront(node)
	return node.e.value, true
}

func (c *LRU[K, V]) unlink(node *node[K, V]) {
	node.previous.next = node.next
	node.next.previous = node.previous
}

func (c *LRU[K, V]) addFront(node *node[K, V]) {
	node.next = c.head.next
	node.previous = c.head
	c.head.next.previous = node
	c.head.next = node
}

func (c *LRU[K, V]) Put(key K, value V) {
	if node, ok := c.cache[key]; ok {
		node.e.value = value
		c.unlink(node)
		c.addFront(node)
		return
	}
	if len(c.cache) >= c.capacity {
		lru := c.tail.previous
		c.unlink(lru)
		delete(c.cache, lru.e.key)
	}
	entry := &entry[K, V]{key: key, value: value}
	node := &node[K, V]{e: entry}
	c.cache[key] = node
	c.addFront(node)
}

func (c *LRU[K, V]) Values() []V {
	values := make([]V, 0, c.Len())
	node := c.head.next
	for node != c.tail {
		values = append(values, node.e.value)
		node = node.next
	}
	return values
}

func (c *LRU[K, V]) Len() int {
	return len(c.cache)
}
