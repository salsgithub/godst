package bidimap

type Map[K, V comparable] struct {
	forward map[K]V
	reverse map[V]K
}

func New[K, V comparable]() *Map[K, V] {
	return &Map[K, V]{
		forward: make(map[K]V),
		reverse: make(map[V]K),
	}
}

func (m *Map[K, V]) Put(key K, value V) {
	if existingValue, ok := m.forward[key]; ok {
		delete(m.reverse, existingValue)
	}
	if existingKey, ok := m.reverse[value]; ok {
		delete(m.forward, existingKey)
	}
	m.forward[key] = value
	m.reverse[value] = key
}

func (m *Map[K, V]) Delete(key K) {
	if value, ok := m.forward[key]; ok {
		delete(m.forward, key)
		delete(m.reverse, value)
	}
}

func (m *Map[K, V]) DeleteValue(value V) {
	if key, ok := m.reverse[value]; ok {
		delete(m.reverse, value)
		delete(m.forward, key)
	}
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	value, ok := m.forward[key]
	return value, ok
}

func (m *Map[K, V]) GetKey(value V) (K, bool) {
	key, ok := m.reverse[value]
	return key, ok
}

func (m *Map[K, V]) Inverse() *Map[V, K] {
	return &Map[V, K]{
		forward: m.reverse,
		reverse: m.forward,
	}
}

func (m *Map[K, V]) Keys() []K {
	size := len(m.forward)
	keys := make([]K, size)
	index := 0
	for key := range m.forward {
		keys[index] = key
		index++
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	size := len(m.reverse)
	values := make([]V, size)
	index := 0
	for value := range m.reverse {
		values[index] = value
		index++
	}
	return values
}

func (m *Map[K, V]) All(yield func(K, V) bool) {
	for k, v := range m.forward {
		if !yield(k, v) {
			return
		}
	}
}

func (m *Map[K, V]) Clear() {
	clear(m.forward)
	clear(m.reverse)
}

func (m *Map[K, V]) IsEmpty() bool {
	return len(m.forward) == 0
}

func (m *Map[K, V]) Len() int {
	return len(m.forward)
}
