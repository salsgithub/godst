package cache

import (
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestLRU_New(t *testing.T) {
	t.Run("lru with capacity <= 0 defaults to 1", func(t *testing.T) {
		l := NewLRU[string, string](0)
		assertx.Equal(t, l.Len(), 0)
		assertx.Equal(t, l.capacity, 1)
	})
	t.Run("lru with small capacity", func(t *testing.T) {
		l := NewLRU[string, string](3)
		assertx.Equal(t, l.Len(), 0)
		assertx.Equal(t, l.capacity, 3)
	})
}

func TestLRU_Get(t *testing.T) {
	t.Run("get an non-existent key", func(t *testing.T) {
		l := NewLRU[string, string](1)
		value, ok := l.Get("language")
		assertx.False(t, ok)
		assertx.Equal(t, value, "")
	})
	t.Run("get returns existing value", func(t *testing.T) {
		l := NewLRU[string, string](1)
		l.Put("A", "B")
		value, ok := l.Get("A")
		assertx.True(t, ok)
		assertx.Equal(t, value, "B")
	})
	t.Run("get marks the value as mru", func(t *testing.T) {
		l := NewLRU[string, string](2)
		l.Put("language", "Go")
		l.Put("front-end", "Svelte")
		assertx.Equal(t, l.Values(), []string{"Svelte", "Go"})
		value, ok := l.Get("language")
		assertx.True(t, ok)
		assertx.Equal(t, value, "Go")
		assertx.Equal(t, l.Values(), []string{"Go", "Svelte"})
	})
}

func TestLRU_Put(t *testing.T) {
	t.Run("put updates value of existing entry", func(t *testing.T) {
		l := NewLRU[string, string](3)
		l.Put("language", "TypeScript")
		l.Put("language", "Go")
		assertx.Equal(t, l.Len(), 1)
	})
	t.Run("put evicts lru when capacity reached", func(t *testing.T) {
		l := NewLRU[string, string](2)
		l.Put("cloud", "GCP")
		l.Put("backend", "Go")
		l.Put("front-end", "Svelte")
		assertx.Equal(t, l.Len(), 2)
		value, ok := l.Get("cloud")
		assertx.False(t, ok)
		assertx.Equal(t, value, "")
	})
}

func TestLRU_Values(t *testing.T) {
	t.Run("values yields values mru to lru in order", func(t *testing.T) {
		l := NewLRU[string, string](2)
		l.Put("cloud", "GCP")
		l.Put("backend", "Go")
		assertx.Equal(t, l.Values(), []string{"Go", "GCP"})
	})
	t.Run("values yields empty for empty lru", func(t *testing.T) {
		l := NewLRU[string, string](2)
		assertx.Equal(t, l.Values(), []string{})
	})
}
