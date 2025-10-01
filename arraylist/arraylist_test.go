package arraylist

import (
	"cmp"
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestArrayList_New(t *testing.T) {
	tests := []struct {
		name             string
		options          []Option[int]
		expectedSize     int
		expectedCapacity int
		expectedElements []int
		expectFixedSize  bool
	}{
		{
			name:             "Empty list",
			options:          nil,
			expectedSize:     0,
			expectedCapacity: 0,
			expectedElements: []int{},
			expectFixedSize:  false,
		},
		{
			name:             "List with empty initial values",
			options:          []Option[int]{WithInitialValues[int]()},
			expectedSize:     0,
			expectedCapacity: 0,
			expectedElements: []int{},
			expectFixedSize:  false,
		},
		{
			name:             "List with initial values",
			options:          []Option[int]{WithInitialValues(1, 2, 3)},
			expectedSize:     3,
			expectedCapacity: 3,
			expectedElements: []int{1, 2, 3},
			expectFixedSize:  false,
		},
		{
			name:             "List with initial capacity",
			options:          []Option[int]{WithInitialCapacity[int](5)},
			expectedSize:     0,
			expectedCapacity: 5,
			expectedElements: []int{},
			expectFixedSize:  false,
		},
		{
			name: "List with initial capacity and values",
			options: []Option[int]{
				WithInitialCapacity[int](100),
				WithInitialValues(7, 21),
			},
			expectedSize:     2,
			expectedCapacity: 100,
			expectedElements: []int{7, 21},
			expectFixedSize:  false,
		},
		{
			name: "List with initial values exceeding fixed size",
			options: []Option[int]{
				WithFixedSize[int](2),
				WithInitialValues(1, 2, 3, 4, 5),
			},
			expectedSize:     2,
			expectedCapacity: 2,
			expectedElements: []int{1, 2},
			expectFixedSize:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(test.options...)
			size := l.Len()
			capacity := cap(l.elements)
			fixedSize := l.IsFixedSize()
			assertx.Equal(t, size, test.expectedSize)
			assertx.Equal(t, capacity, test.expectedCapacity)
			assertx.Equal(t, l.elements, test.expectedElements)
			assertx.Equal(t, fixedSize, test.expectFixedSize)
		})
	}
}

func TestArrayList_OrderOfOptions(t *testing.T) {
	l := New(WithInitialValues(1, 2, 3), WithInitialCapacity[int](25))
	assertx.Equal(t, l.Len(), 3)
	assertx.Equal(t, cap(l.elements), 25)
}

func TestArrayList_DynamicSizing(t *testing.T) {
	tests := []struct {
		name             string
		options          []Option[int]
		actions          func(*List[int])
		expectedSize     int
		expectedCapacity int
	}{
		{
			name:    "List with default growth factor",
			options: []Option[int]{WithInitialCapacity[int](2)},
			actions: func(l *List[int]) {
				l.Add(1, 2)
				l.Add(3)
			},
			expectedSize:     3,
			expectedCapacity: 4,
		},
		{
			name: "List with invalid growth factor",
			options: []Option[int]{
				WithInitialCapacity[int](2),
				WithGrowthFactor[int](1), // grow default back to 2.0
			},
			actions: func(l *List[int]) {
				l.Add(1, 2)
				l.Add(3)
			},
			expectedSize:     3,
			expectedCapacity: 4,
		},
		{
			name: "List grows with respect to growth factor",
			options: []Option[int]{
				WithInitialCapacity[int](10),
				WithGrowthFactor[int](1.5), // grow by 1.5x
			},
			actions: func(l *List[int]) {
				for i := range 10 {
					l.Add(i)
				}
				l.Add(11)
			},
			expectedSize:     11,
			expectedCapacity: 15,
		},
		{
			name: "List with default shrink factor",
			options: []Option[int]{
				WithInitialCapacity[int](2),
			},
			actions: func(l *List[int]) {
				for i := range 10 {
					l.Add(i)
				}
				for i := range 5 {
					l.Remove(i)
				}
			},
			expectedSize:     5,
			expectedCapacity: 16,
		},
		{
			name: "List does not shrink below threshold",
			options: []Option[int]{
				WithInitialCapacity[int](10),
			},
			actions: func(l *List[int]) {
				for i := range 10 {
					l.Add(i)
				}
				// 20% full after removing 8 elements
				for range 8 {
					l.Remove(0)
				}
			},
			expectedSize:     2,
			expectedCapacity: 2, // shrink to fit exact size
		},
		{
			name: "List shrinks with respect to shrink factor",
			options: []Option[int]{
				WithInitialCapacity[int](10),
				WithShrinkFactor[int](0.8), // shrink when less than 80% full
			},
			actions: func(l *List[int]) {
				for i := range 10 {
					l.Add(i)
				}
				for range 3 {
					l.Remove(0)
				}
			},
			expectedSize:     7,
			expectedCapacity: 7,
		},
		{
			name: "List with disabled shrink",
			options: []Option[int]{
				WithShrinkFactor[int](0.0),
			},
			actions: func(l *List[int]) {
				for i := range 10 {
					l.Add(i)
				}
				for range 5 {
					l.Remove(0)
				}
			},
			expectedSize:     5,
			expectedCapacity: 16,
		},
		{
			name: "List growth is capped by fixed size",
			options: []Option[int]{
				WithInitialCapacity[int](2),
				WithGrowthFactor[int](10.0),
				WithFixedSize[int](5),
			},
			actions: func(l *List[int]) {
				l.Add(1, 2)
				l.Add(3)
			},
			expectedSize:     3,
			expectedCapacity: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(test.options...)
			test.actions(l)
			size := l.Len()
			capacity := cap(l.elements)
			assertx.Equal(t, size, test.expectedSize)
			assertx.Equal(t, capacity, test.expectedCapacity)
		})
	}
}

func TestArrayList_Add(t *testing.T) {
	l := New[int]()
	l.Add(1, 2, 3)
	assertx.Equal(t, l.elements, []int{1, 2, 3})
	t.Run("adding nothing has no changes", func(t *testing.T) {
		l.Add()
		assertx.Equal(t, l.elements, []int{1, 2, 3})
	})
	t.Run("adding to fixed size list when limit reached does nothing", func(t *testing.T) {
		frozen := New(WithInitialValues(1, 2), WithFixedSize[int](2))
		frozen.Add(1, 2, 3, 4, 5)
		assertx.Equal(t, frozen.Len(), 2)
	})
}

func TestArrayList_CheckBounds(t *testing.T) {
	l := New[int]()
	l.Add(1)
	assertx.NotNil(t, l.checkBounds(21))
}

func TestArrayList_Remove(t *testing.T) {
	t.Run("remove at valid index", func(t *testing.T) {
		l := New(WithInitialValues(1, 2, 3, 4, 5))
		assertx.Nil(t, l.Remove(3))
		assertx.Equal(t, l.Len(), 4)
		assertx.Equal(t, l.elements, []int{1, 2, 3, 5})
	})
	t.Run("remove at invalid index errors", func(t *testing.T) {
		l := New(WithInitialValues(1, 2))
		assertx.ErrorIs(t, l.Remove(2), ErrIndexOutOfBounds)
	})
}

func TestArrayList_Replace(t *testing.T) {
	l := New(WithInitialValues("house", "table", "chair"))
	t.Run("replace at valid index", func(t *testing.T) {
		assertx.Nil(t, l.Replace(2, "sofa"))
		assertx.Equal(t, l.elements, []string{"house", "table", "sofa"})
	})
	t.Run("should error if replace is invalid", func(t *testing.T) {
		assertx.ErrorIs(t, l.Replace(10, "a"), ErrIndexOutOfBounds)
	})
}

func TestArrayList_Value(t *testing.T) {
	l := New(WithInitialValues(1, 2))
	t.Run("retrieve first value", func(t *testing.T) {
		value, err := l.Value(0)
		assertx.Nil(t, err)
		assertx.Equal(t, value, 1)
	})
	t.Run("retrieve value from invalid index", func(t *testing.T) {
		_, err := l.Value(20)
		assertx.ErrorIs(t, err, ErrIndexOutOfBounds)
	})
}

func TestArrayList_Values(t *testing.T) {
	l := New(WithInitialValues(1, 2))
	assertx.Equal(t, l.Values(), []int{1, 2})
}

func TestArrayList_Contains(t *testing.T) {
	l := New(WithInitialValues(1, 2))
	assertx.True(t, l.Contains(1))
}

func TestArrayList_ContainsAll(t *testing.T) {
	l := New(WithInitialValues(1, 2, 3))
	assertx.True(t, l.ContainsAll())
	assertx.True(t, l.ContainsAll(1, 2))
	t.Run("list does not contain all elements", func(t *testing.T) {
		assertx.False(t, l.ContainsAll(1, 5))
	})
}

func TestArrayList_Sort(t *testing.T) {
	l := New(WithInitialValues("c", "a", "b"))
	l.Sort(cmp.Compare[string])
	assertx.Equal(t, l.Values(), []string{"a", "b", "c"})
}

func TestArrayList_Clear(t *testing.T) {
	l := New[int]()
	for i := range 10 {
		l.Add(i)
	}
	initialCapacity := cap(l.elements)
	l.Clear()
	assertx.Equal(t, l.Len(), 0)
	assertx.Equal(t, cap(l.elements), initialCapacity)
}

func TestArrayList_Reset(t *testing.T) {
	initialCapacity := 7
	l := New(WithInitialValues("a", "b", "c", "d"), WithInitialCapacity[string](initialCapacity))
	l.Reset()
	assertx.Equal(t, l.Len(), 0)
	assertx.Equal(t, cap(l.elements), initialCapacity)
}

func TestArrayList_Reverse(t *testing.T) {
	l := New(WithInitialValues("table", "chair", "bed"))
	l.Reverse()
	assertx.Equal(t, l.elements, []string{"bed", "chair", "table"})
}

func TestArrayList_IsEmpty(t *testing.T) {
	l := New[int]()
	assertx.True(t, l.IsEmpty())
	l.Add(0)
	assertx.False(t, l.IsEmpty())
}

func TestArrayList_ManualGrow(t *testing.T) {
	l := New(WithInitialCapacity[int](10))
	l.Add(1, 2, 3)
	underlyingBefore := &l.elements[0]
	l.grow(5)
	assertx.Equal(t, cap(l.elements), 10)
	underlyingAfter := &l.elements[0]
	assertx.Equal(t, underlyingBefore, underlyingAfter)
}

func TestArrayList_Shrink(t *testing.T) {
	l := New[int]()
	l.shrink()
	assertx.Equal(t, l.Len(), 0)
	assertx.Equal(t, cap(l.elements), 0)
}

func TestArrayList_String(t *testing.T) {
	cars := New[string]()
	cars.Add("Audi", "Porsche")
	assertx.Equal(t, cars.String(), "[Audi Porsche]")
}

func BenchmarkAdd_100(b *testing.B) {
	size := 100
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_1_000(b *testing.B) {
	size := 1_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_10_000(b *testing.B) {
	size := 10_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_100_000(b *testing.B) {
	size := 100_000
	for b.Loop() {
		l := New[int]()
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_100_WithInitialCapacity(b *testing.B) {
	size := 100
	for b.Loop() {
		l := New(WithInitialCapacity[int](size))
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_1_000_WithInitialCapacity(b *testing.B) {
	size := 1_000
	for b.Loop() {
		l := New(WithInitialCapacity[int](size))
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_10_000_WithInitialCapacity(b *testing.B) {
	size := 10_000
	for b.Loop() {
		l := New(WithInitialCapacity[int](size))
		for i := range size {
			l.Add(i)
		}
	}
}

func BenchmarkAdd_100_000_WithInitialCapacity(b *testing.B) {
	size := 100_000
	for b.Loop() {
		l := New(WithInitialCapacity[int](size))
		for i := range size {
			l.Add(i)
		}
	}
}
