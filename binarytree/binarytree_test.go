package binarysearchtree

import (
	"math/rand/v2"
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestBinarySearchTree_New(t *testing.T) {
	tree := New(10, 20, 30)
	assertx.NotNil(t, tree)
}

func TestBinarySearchTree_Insert(t *testing.T) {
	t.Run("insert single value", func(t *testing.T) {
		tree := New[int]()
		tree.Insert(10)
		assertx.Equal(t, tree.root, &node[int]{value: 10})
		assertx.Equal(t, tree.Len(), 1)
	})
	t.Run("insert multiple values", func(t *testing.T) {
		tree := New(10, 9, 11)
		left := &node[int]{
			value: 9,
			left:  nil,
			right: nil,
		}
		right := &node[int]{
			value: 11,
			left:  nil,
			right: nil,
		}
		root := &node[int]{
			value: 10,
			left:  left,
			right: right,
		}
		assertx.Equal(t, tree.root, root)
		assertx.Equal(t, tree.Len(), 3)
		assertx.True(t, tree.Contains(9))
		assertx.True(t, tree.Contains(10))
		assertx.True(t, tree.Contains(11))
	})
}

func TestBinarySearchTree_InsertAll(t *testing.T) {
	tree := New[string]()
	tree.InsertAll("Mango", "Apple", "Strawberry", "Blueberry", "Pineapple")
	assertx.Equal(t, tree.Len(), 5)
}

func TestBinarySearchTree_Remove(t *testing.T) {
	t.Run("remove from empty tree", func(t *testing.T) {
		tree := New[int]()
		tree.Remove(0)
		assertx.Equal(t, tree.size, 0)
	})
	t.Run("remove a leaf from the right", func(t *testing.T) {
		tree := New(10, 9, 11)
		tree.Remove(11)
		assertx.Equal(t, tree.Len(), 2)
		assertx.False(t, tree.Contains(11))
	})
	t.Run("remove a leaf from the left", func(t *testing.T) {
		tree := New(10, 9, 11)
		tree.Remove(9)
		assertx.Equal(t, tree.Len(), 2)
		assertx.False(t, tree.Contains(9))
	})
	t.Run("remove a node with only a right child", func(t *testing.T) {
		tree := New(10, 5, 7)
		tree.Remove(5)
		assertx.Equal(t, tree.Len(), 2)
		assertx.False(t, tree.Contains(5))
	})
	t.Run("remove a node with only a left child", func(t *testing.T) {
		tree := New(10, 15, 12)
		tree.Remove(15)
		assertx.Equal(t, tree.Len(), 2)
		assertx.False(t, tree.Contains(15))
	})
	t.Run("remove root node with both children", func(t *testing.T) {
		tree := New(10, 5, 15)
		tree.Remove(10)
		assertx.Equal(t, tree.Len(), 2)
		assertx.False(t, tree.Contains(10))
	})
	t.Run("remove internal node with both children", func(t *testing.T) {
		tree := New(10, 5, 15, 3, 7)
		tree.Remove(5)
		assertx.Equal(t, tree.Len(), 4)
		assertx.False(t, tree.Contains(5))
	})
}

func TestBinarySearchTree_RemoveAll(t *testing.T) {
	tree := New(1, 2, 3)
	tree.RemoveAll(1, 2, 3)
	assertx.Equal(t, tree.Len(), 0)
	assertx.True(t, tree.IsEmpty())
}

func TestBinarySearchTree_MinNode(t *testing.T) {
	tree := New(10, 2, 3)
	minNode := minNode(tree.root)
	assertx.NotNil(t, minNode)
	assertx.Equal(t, minNode.value, 2)
}

func TestBinarySearchTree_Height(t *testing.T) {
	t.Run("empty tree has height 0", func(t *testing.T) {
		tree := New[int]()
		assertx.Equal(t, tree.Height(), 0)
	})
	t.Run("tree with just a root has height 1", func(t *testing.T) {
		tree := New[string]()
		tree.Insert("Go")
		assertx.Equal(t, tree.Height(), 1)
	})
	t.Run("tree with max height on the left", func(t *testing.T) {
		tree := New("Z", "A", "B", "C", "D", "E", "F")
		assertx.Equal(t, tree.Height(), 7)
	})
	t.Run("tree with max height on the right", func(t *testing.T) {
		tree := New("Go", "A", "H", "I", "J", "K", "L")
		assertx.Equal(t, tree.Height(), 6)
	})
}

func TestBinarySearchTree_Min(t *testing.T) {
	t.Run("min on empty tree yields zero value", func(t *testing.T) {
		tree := New[int]()
		value, ok := tree.Min()
		assertx.False(t, ok)
		assertx.Equal(t, value, 0)
	})
	t.Run("min on tree yields minimum value", func(t *testing.T) {
		tree := New(2009, 1971, 2020)
		value, ok := tree.Min()
		assertx.True(t, ok)
		assertx.Equal(t, value, 1971)
	})
}

func TestBinarySearchTree_Max(t *testing.T) {
	t.Run("max on empty tree yields zero value", func(t *testing.T) {
		tree := New[int]()
		value, ok := tree.Max()
		assertx.False(t, ok)
		assertx.Equal(t, value, 0)
	})
	t.Run("max on tree yields maximum value", func(t *testing.T) {
		tree := New(2009, 1971, 2020)
		value, ok := tree.Max()
		assertx.True(t, ok)
		assertx.Equal(t, value, 2020)
	})
}

func TestBinarySearchTree_Contains(t *testing.T) {
	tree := New(10)
	assertx.True(t, tree.Contains(10))
	assertx.False(t, tree.Contains(11))
}

func TestBinarySearchTree_Traversals(t *testing.T) {
	tree := New(10, 5, 15, 3, 7, 12, 18)
	/*
				10
			    / \
		       5   15
		      / \  / \
			 3  7 12 18
	*/
	t.Run("pre-order traversal", func(t *testing.T) {
		assertx.Equal(t, tree.ValuesPreOrder(), []int{10, 5, 3, 7, 15, 12, 18})
	})
	t.Run("in-order traversal", func(t *testing.T) {
		assertx.Equal(t, tree.ValuesInOrder(), []int{3, 5, 7, 10, 12, 15, 18})
	})
	t.Run("post-order traversal", func(t *testing.T) {
		assertx.Equal(t, tree.ValuesPostOrder(), []int{3, 7, 5, 12, 18, 15, 10})
	})
	t.Run("breadth first search on empty tree yields empty", func(t *testing.T) {
		tree := New[int]()
		assertx.Equal(t, tree.ValuesBreadthFirst(), []int{})
	})
	t.Run("breadth first search", func(t *testing.T) {
		assertx.Equal(t, tree.ValuesBreadthFirst(), []int{10, 5, 15, 3, 7, 12, 18})
	})
}

func BenchmarkInsertWorstCase_100(b *testing.B) {
	size := 100
	for b.Loop() {
		t := New[int]()
		for i := range size {
			t.Insert(i)
		}
	}
}

func BenchmarkInsertWorstCase_1_000(b *testing.B) {
	size := 1_000
	for b.Loop() {
		t := New[int]()
		for i := range size {
			t.Insert(i)
		}
	}
}

func BenchmarkInsertWorstCase_10_000(b *testing.B) {
	size := 10_000
	for b.Loop() {
		t := New[int]()
		for i := range size {
			t.Insert(i)
		}
	}
}

func BenchmarkInsertAverageCase_100(b *testing.B) {
	size := 100
	permutation := rand.Perm(size)
	for b.Loop() {
		t := New[int]()
		t.InsertAll(permutation...)
	}
}

func BenchmarkInsertAverageCase_1_000(b *testing.B) {
	size := 1_000
	permutation := rand.Perm(size)
	for b.Loop() {
		t := New[int]()
		t.InsertAll(permutation...)
	}
}

func BenchmarkInsertAverageCase_10_000(b *testing.B) {
	size := 10_000
	permutation := rand.Perm(size)
	for b.Loop() {
		t := New[int]()
		t.InsertAll(permutation...)
	}
}
