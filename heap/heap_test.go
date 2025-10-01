package heap

import (
	"math/rand/v2"
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestHeap_New(t *testing.T) {
	h := New(func(a, b int) bool {
		return a > b
	})
	assertx.NotNil(t, h)
}

func TestHeap_NewMin(t *testing.T) {
	h := NewMin[int]()
	assertx.NotNil(t, h)
}

func TestHeap_NewMax(t *testing.T) {
	h := NewMax[int]()
	assertx.NotNil(t, h)
}

func TestHeap_Min(t *testing.T) {
	h := NewMin[int]()
	h.Push(10)
	h.Push(7)
	h.Push(21)
	h.Push(99)
	assertx.Equal(t, h.Len(), 4)
	top, ok := h.Peek()
	assertx.True(t, ok)
	assertx.Equal(t, top, 7)
	pop, ok := h.Pop()
	assertx.True(t, ok)
	assertx.Equal(t, pop, 7)
	assertx.Equal(t, h.Len(), 3)
}

func TestHeap_Max(t *testing.T) {
	h := NewMax[int]()
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	assertx.Equal(t, h.Len(), 4)
	top, ok := h.Peek()
	assertx.True(t, ok)
	assertx.Equal(t, top, 4)
	pop, ok := h.Pop()
	assertx.True(t, ok)
	assertx.Equal(t, pop, 4)
	assertx.Equal(t, h.Len(), 3)
}

func TestHeap_MinPop(t *testing.T) {
	t.Run("pop from empty heap yields zero value", func(t *testing.T) {
		h := NewMin[int]()
		pop, ok := h.Pop()
		assertx.False(t, ok)
		assertx.Equal(t, pop, 0)
	})
	t.Run("pop from min heap returns values in increasing order", func(t *testing.T) {
		h := NewMin[int]()
		h.Push(20)
		h.Push(50)
		h.Push(88)
		h.Push(32)
		values := []int{}
		for !h.IsEmpty() {
			pop, ok := h.Pop()
			assertx.True(t, ok)
			values = append(values, pop)
		}
		assertx.Equal(t, h.Len(), 0)
		assertx.Equal(t, values, []int{20, 32, 50, 88})
	})
	t.Run("pop from min heap triggers swap with right child", func(t *testing.T) {
		h := NewMin[int]()
		h.Push(10)
		h.Push(20)
		h.Push(5)
		h.Push(30)
		pop, ok := h.Pop()
		assertx.True(t, ok)
		assertx.Equal(t, pop, 5)
		peek, ok := h.Peek()
		assertx.True(t, ok)
		assertx.Equal(t, peek, 10)
	})
}

func TestHeap_MaxPop(t *testing.T) {
	t.Run("pop from empty heap yields zero value", func(t *testing.T) {
		h := NewMax[int]()
		pop, ok := h.Pop()
		assertx.False(t, ok)
		assertx.Equal(t, pop, 0)
	})
	t.Run("pop from max heap returns values in decreasing order", func(t *testing.T) {
		h := NewMax[int]()
		h.Push(1)
		h.Push(2)
		h.Push(3)
		h.Push(4)
		values := []int{}
		for !h.IsEmpty() {
			pop, ok := h.Pop()
			assertx.True(t, ok)
			values = append(values, pop)
		}
		assertx.Equal(t, h.Len(), 0)
		assertx.Equal(t, values, []int{4, 3, 2, 1})
	})
}

func TestHeap_Peek(t *testing.T) {
	t.Run("peek on empty heap yields zero value", func(t *testing.T) {
		h := NewMax[int]()
		peek, ok := h.Peek()
		assertx.False(t, ok)
		assertx.Equal(t, peek, 0)
	})
	t.Run("peek on min heap yields min value without removing it", func(t *testing.T) {
		h := NewMin[int]()
		h.Push(7)
		h.Push(1)
		h.Push(25)
		peek, ok := h.Peek()
		assertx.True(t, ok)
		assertx.Equal(t, peek, 1)
		assertx.Equal(t, h.Len(), 3)
	})
	t.Run("peek on max heap yields max value without removing it", func(t *testing.T) {
		h := NewMax[int]()
		h.Push(7)
		h.Push(1)
		h.Push(25)
		peek, ok := h.Peek()
		assertx.True(t, ok)
		assertx.Equal(t, peek, 25)
		assertx.Equal(t, h.Len(), 3)
	})
}

func TestHeap_Values(t *testing.T) {
	h := NewMax[int]()
	h.Push(7)
	h.Push(1)
	h.Push(25)
	assertx.Equal(t, h.Values(), []int{25, 1, 7})
}

func TestHeap_Clear(t *testing.T) {
	h := NewMax[int]()
	for i := range 20 {
		h.Push(i)
	}
	h.Clear()
	assertx.Equal(t, h.Len(), 0)
	assertx.True(t, h.IsEmpty())
}

func BenchmarkHeap_Push_100(b *testing.B) {
	size := 100
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
	}
}

func BenchmarkHeap_Push_1_000(b *testing.B) {
	size := 1_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
	}
}

func BenchmarkHeap_Push_10_000(b *testing.B) {
	size := 10_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
	}
}

func BenchmarkHeap_Push_100_000(b *testing.B) {
	size := 100_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
	}
}

func BenchmarkHeap_Pop_100(b *testing.B) {
	size := 100
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
		for !h.IsEmpty() {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop_1_000(b *testing.B) {
	size := 1_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
		for !h.IsEmpty() {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop_10_000(b *testing.B) {
	size := 10_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
		for !h.IsEmpty() {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop_100_000(b *testing.B) {
	size := 100_000
	permutation := rand.Perm(size)
	b.ResetTimer()
	for b.Loop() {
		h := NewMin[int]()
		for _, value := range permutation {
			h.Push(value)
		}
		for !h.IsEmpty() {
			h.Pop()
		}
	}
}
