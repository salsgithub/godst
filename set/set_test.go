package set

import (
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestSet_New(t *testing.T) {
	s := New("language", "Go", "Go")
	assertx.Equal(t, s.Len(), 2)
	assertx.False(t, s.IsEmpty())
}

func TestSet_Remove(t *testing.T) {
	s := New("Tony", "Stark")
	s.Remove("Tony")
	assertx.Equal(t, s.Len(), 1)
	assertx.Equal(t, s.Values(), []string{"Stark"})
	assertx.False(t, s.Contains("Tony"))
}

func TestSet_Union(t *testing.T) {
	s := New("Tony", "Stark")
	a := New("Bruce", "Banner")
	union := s.Union(a)
	assertx.Equal(t, union.Len(), 4)
	assertx.True(t, union.Contains("Tony"))
	assertx.True(t, union.Contains("Stark"))
	assertx.True(t, union.Contains("Bruce"))
	assertx.True(t, union.Contains("Banner"))
}

func TestSet_Intersection(t *testing.T) {
	t.Run("intersection with larger set", func(t *testing.T) {
		s := New("Spiderman", "Iron Man", "Hulk")
		a := New("Spiderman", "Iron Man")
		intersection := s.Intersection(a)
		assertx.True(t, intersection.Contains("Spiderman"))
		assertx.True(t, intersection.Contains("Iron Man"))
	})
	t.Run("intersection with smaller set", func(t *testing.T) {
		s := New("Spiderman", "Iron Man", "Hulk")
		a := New("Spiderman", "Iron Man")
		intersection := a.Intersection(s)
		assertx.Equal(t, intersection.Len(), 2)
		assertx.True(t, intersection.Contains("Spiderman"))
		assertx.True(t, intersection.Contains("Iron Man"))
	})
}

func TestSet_Delta(t *testing.T) {
	s := New("London", "New York", "Brazil", "Portugal")
	a := New("London", "New York")
	delta := s.Delta(a)
	assertx.Equal(t, delta.Len(), 2)
	assertx.True(t, delta.Contains("Brazil"))
	assertx.True(t, delta.Contains("Portugal"))
}

func TestSet_All(t *testing.T) {
	t.Run("iterates all values", func(t *testing.T) {
		s := New("London", "New York")
		results := make(map[string]bool)
		for value := range s.All {
			results[value] = true
		}
		assertx.True(t, results["London"])
		assertx.True(t, results["New York"])
	})
	t.Run("iteration stops when yield fails", func(t *testing.T) {
		s := New(10, 20, 30, 40)
		count := 0
		s.All(func(i int) bool {
			count++
			return count < 2
		})
		assertx.Equal(t, count, 2)
	})
}

func TestSet_Clear(t *testing.T) {
	s := New("London", "New York")
	s.Clear()
	assertx.Equal(t, s.Len(), 0)
	assertx.True(t, s.IsEmpty())
}

func TestSet_String(t *testing.T) {
	t.Run("string set return sorted strings for string", func(t *testing.T) {
		s := New("London", "New York", "Amsterdam")
		assertx.Equal(t, s.String(), "Amsterdam, London, New York")
	})
	t.Run("int set returns sorted ints for string", func(t *testing.T) {
		s := New(19, 71, 7, 2009)
		assertx.Equal(t, s.String(), "7, 19, 71, 2009")
	})
	t.Run("float64 set returns sorted floats for string", func(t *testing.T) {
		s := New(19.71, 20.09)
		assertx.Equal(t, s.String(), "19.71, 20.09")
	})
}
