package bidimap

import (
	"sort"
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestBidiMap_New(t *testing.T) {
	m := New[int, int]()
	assertx.NotNil(t, m)
	assertx.Equal(t, m.Len(), 0)
}

func TestBidiMap_Put(t *testing.T) {
	t.Run("put overrides any existing value", func(t *testing.T) {
		m := New[string, int]()
		m.Put("Go", 99)
		assertx.Equal(t, m.Len(), 1)
		assertx.Equal(t, len(m.reverse), 1)
		m.Put("Go", 95)
		assertx.Equal(t, m.Len(), 1)
	})
	t.Run("put overrides any existing key and value", func(t *testing.T) {
		m := New[string, int]()
		m.Put("A", 1)
		m.Put("B", 2)
		m.Put("A", 2)
		assertx.Equal(t, m.Len(), 1)
		assertx.Equal(t, len(m.forward), 1)
		assertx.Equal(t, len(m.reverse), 1)
	})
}

func TestBidiMap_Delete(t *testing.T) {
	t.Run("delete empty map does nothing", func(t *testing.T) {
		m := New[string, string]()
		assertx.Equal(t, m.Len(), 0)
		m.Delete("something")
		assertx.Equal(t, m.Len(), 0)
	})
	t.Run("delete key from map", func(t *testing.T) {
		m := New[string, string]()
		m.Put("hello", "there")
		assertx.Equal(t, m.Len(), 1)
		m.Delete("hello")
		assertx.Equal(t, m.Len(), 0)
	})
}

func TestBidiMap_DeleteValue(t *testing.T) {
	t.Run("delete value from empty map does nothing", func(t *testing.T) {
		m := New[string, string]()
		m.DeleteValue("not there")
		assertx.Equal(t, m.Len(), 0)
	})

	t.Run("delete value from map", func(t *testing.T) {
		m := New[string, string]()
		m.Put("a", "b")
		m.Put("c", "d")
		assertx.Equal(t, m.Len(), 2)
		m.DeleteValue("d")
		assertx.Equal(t, m.Len(), 1)
	})
}

func TestBidiMap_Get(t *testing.T) {
	t.Run("finding non-existent key", func(t *testing.T) {
		m := New[string, int]()
		m.Put("Go", 99)
		value, ok := m.Get("G")
		assertx.False(t, ok)
		assertx.Equal(t, value, 0)
	})
	t.Run("finding existing key", func(t *testing.T) {
		m := New[string, int]()
		m.Put("Go", 99)
		value, ok := m.Get("Go")
		assertx.True(t, ok)
		assertx.Equal(t, value, 99)
	})
}

func TestBidiMap_GetKey(t *testing.T) {
	t.Run("finding non-existent value", func(t *testing.T) {
		m := New[string, int]()
		m.Put("Go", 99)
		value, ok := m.GetKey(0)
		assertx.False(t, ok)
		assertx.Equal(t, value, "")
	})
	t.Run("finding existing value", func(t *testing.T) {
		m := New[string, int]()
		m.Put("Go", 99)
		value, ok := m.GetKey(99)
		assertx.True(t, ok)
		assertx.Equal(t, value, "Go")
	})
}

func TestBidiMap_Inverse(t *testing.T) {
	m := New[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	inverse := m.Inverse()
	assertx.Equal(t, m.Len(), inverse.Len())
	one, ok := inverse.Get(1)
	assertx.True(t, ok)
	assertx.Equal(t, one, "a")
	two, ok := inverse.Get(2)
	assertx.True(t, ok)
	assertx.Equal(t, two, "b")
}

func TestBidiMap_Keys(t *testing.T) {
	t.Run("empty map yields no keys", func(t *testing.T) {
		m := New[string, string]()
		assertx.Equal(t, m.Keys(), []string{})
	})
	t.Run("populated map yields keys", func(t *testing.T) {
		m := New[string, string]()
		m.Put("a", "b")
		m.Put("c", "d")
		keys := m.Keys()
		sort.Strings(keys)
		assertx.Equal(t, keys, []string{"a", "c"})
	})
}

func TestBidiMap_Values(t *testing.T) {
	t.Run("empty map yields no values", func(t *testing.T) {
		m := New[string, string]()
		assertx.Equal(t, m.Values(), []string{})
	})
	t.Run("populated map yields values", func(t *testing.T) {
		m := New[string, string]()
		m.Put("a", "b")
		m.Put("c", "d")
		values := m.Values()
		sort.Strings(values)
		assertx.Equal(t, values, []string{"b", "d"})
	})
}

func TestBidiMap_All(t *testing.T) {
	t.Run("iterate an empty map yields empty", func(t *testing.T) {
		m := New[string, string]()
		kvMap := make(map[string]string)
		for k, v := range m.All {
			kvMap[k] = v
		}
		assertx.Equal(t, len(kvMap), 0)
	})
	t.Run("iterate a populated map yields key value pairs", func(t *testing.T) {
		m := New[string, string]()
		m.Put("language", "Go")
		kvMap := make(map[string]string)
		for k, v := range m.All {
			kvMap[k] = v
		}
		assertx.Equal(t, len(kvMap), 1)
	})
	t.Run("iterate using all directly", func(t *testing.T) {
		m := New[string, string]()
		m.Put("language", "Go")
		m.Put("other", "C++")
		kvMap := make(map[string]string)
		m.All(func(key, value string) bool {
			if len(value) < 3 {
				kvMap[key] = value
				return true
			}
			return false
		})
		assertx.Equal(t, len(kvMap), 1)
	})
}

func TestBidiMap_Clear(t *testing.T) {
	m := New[string, string]()
	m.Put("a", "b")
	m.Put("c", "d")
	m.Clear()
	assertx.Equal(t, m.Len(), 0)
	assertx.True(t, m.IsEmpty())
}
