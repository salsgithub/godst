package graph

import (
	"math/rand/v2"
	"testing"

	"github.com/salsgithub/godst/assertx"
)

func TestGraph_AddNode(t *testing.T) {
	g := New[string]()
	g.AddNode("London")
	g.AddNode("London")
	assertx.Equal(t, g.Len(), 1)
}

func TestGraph_AddEdge(t *testing.T) {
	g := New[string]()
	g.AddEdge("London", "New York", 550)
	assertx.NotNil(t, g)
	assertx.Equal(t, g.Len(), 2)
}

func TestGraph_Nodes(t *testing.T) {
	t.Run("nodes returns sorted values for ints", func(t *testing.T) {
		g := New[int]()
		g.AddEdge(1, 2, 3)
		g.AddEdge(4, 5, 6)
		g.AddEdge(7, 8, 9)
		assertx.Equal(t, g.Nodes(), []int{1, 2, 4, 5, 7, 8})
	})

	t.Run("nodes returns sorted values for strings", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("apple", "banana", 3)
		g.AddEdge("blueberry", "strawberry", 7)
		g.AddEdge("carrot", "dragonfruit", 6)
		assertx.Equal(t, g.Nodes(), []string{"apple", "banana", "blueberry", "carrot", "dragonfruit", "strawberry"})
	})
	t.Run("nodes returns sorted values for float64", func(t *testing.T) {
		g := New[float64]()
		g.AddEdge(10, 20, 30)
		g.AddEdge(0.1, 0.2, 1)
		g.AddEdge(19.71, 20.09, 26)
		assertx.Equal(t, g.Nodes(), []float64{0.1, 0.2, 10, 19.71, 20, 20.09})
	})
}

func TestGraph_DeleteNode(t *testing.T) {
	t.Run("deleting non existent node does nothing", func(t *testing.T) {
		g := New[string]()
		g.AddNode("London")
		g.DeleteNode("France")
		assertx.Equal(t, g.Len(), 1)
	})
	t.Run("deleting a node removes it and all connected edges", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 10)
		g.AddEdge("B", "C", 20)
		g.AddEdge("C", "A", 30)
		g.DeleteNode("B")
		assertx.Equal(t, g.Len(), 2)
		expected := map[string][]Edge[string]{
			"A": {},
			"C": {
				Edge[string]{
					Link:   "A",
					Weight: 30,
				},
			},
		}
		assertx.Equal(t, g.adjacency, expected)
	})
}

func TestGraph_Neighbours(t *testing.T) {
	g := New[string]()
	g.AddEdge("Sweden", "Norway", 1)
	g.AddEdge("Spain", "France", 2)
	g.AddEdge("Spain", "Portugal", 3)
	neighbours, ok := g.Neighbours("Spain")
	assertx.True(t, ok)
	assertx.Equal(t, neighbours, []Edge[string]{
		{
			Link:   "France",
			Weight: 2,
		},
		{
			Link:   "Portugal",
			Weight: 3,
		},
	})
}

func TestGraph_HasCycle(t *testing.T) {
	t.Run("empty graph has no cycle", func(t *testing.T) {
		g := New[string]()
		assertx.False(t, g.HasCycle())
	})
	t.Run("graph with one node has no cycle", func(t *testing.T) {
		g := New[string]()
		g.AddNode("Go")
		assertx.False(t, g.HasCycle())
	})
	t.Run("graph with self referencing edge has cycle", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("Go", "Go", 0)
		assertx.True(t, g.HasCycle())
	})
	t.Run("graph with no cycles (DAG)", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("C", "D", 0)
		g.AddEdge("E", "F", 0)
		assertx.False(t, g.HasCycle())
	})
	t.Run("graph with direct cycle", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("B", "A", 0)
		assertx.True(t, g.HasCycle())
	})
	t.Run("graph with long cycle", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("B", "C", 0)
		g.AddEdge("C", "D", 0)
		g.AddEdge("D", "A", 0)
		assertx.True(t, g.HasCycle())
	})
	t.Run("disconnected graph has cycle in sub graph", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("X", "Y", 0)
		g.AddEdge("Y", "Z", 0)
		g.AddEdge("Z", "X", 0)
		assertx.True(t, g.HasCycle())
	})
	t.Run("backtracking cycle", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("A", "C", 0)
		g.AddEdge("C", "D", 0)
		g.AddEdge("D", "A", 0)
		assertx.True(t, g.HasCycle())
	})
}

func TestGraph_TopologicalSort(t *testing.T) {
	t.Run("returns error for graph with cycle", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("B", "A", 0)
		sorted, err := g.TopologicalSort()
		assertx.NotNil(t, err)
		assertx.Nil(t, sorted)
	})
	t.Run("topological sort for DAG", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 0)
		g.AddEdge("B", "C", 0)
		g.AddEdge("C", "D", 0)
		sorted, err := g.TopologicalSort()
		assertx.Nil(t, err)
		assertx.Equal(t, sorted, []string{"A", "B", "C", "D"})
	})
}

func TestGraph_BFS(t *testing.T) {
	t.Run("bfs on empty graph yields error", func(t *testing.T) {
		g := New[string]()
		err := g.BFS("", nil)
		assertx.NotNil(t, err)
	})
	t.Run("bfs with valid start node", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 1)
		g.AddEdge("A", "C", 1)
		g.AddEdge("B", "D", 1)
		g.AddEdge("B", "E", 1)
		g.AddEdge("C", "F", 1)
		g.AddEdge("E", "F", 1)
		/*
				A
			   / \
			  B   C
			 / \   \
			D   E - F
		*/
		order := make([]string, 0)
		onVisit := func(node string) {
			order = append(order, node)
		}
		err := g.BFS("A", onVisit)
		assertx.Nil(t, err)
		assertx.Equal(t, order, []string{"A", "B", "C", "D", "E", "F"})
	})
}

func TestGraph_DFS(t *testing.T) {
	t.Run("dfs on empty graph yields error", func(t *testing.T) {
		g := New[string]()
		err := g.DFS("", nil)
		assertx.NotNil(t, err)
	})
	t.Run("dfs with valid start node", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 1)
		g.AddEdge("A", "C", 1)
		g.AddEdge("B", "D", 1)
		g.AddEdge("B", "E", 1)
		g.AddEdge("C", "F", 1)
		g.AddEdge("E", "F", 1)
		/*
				A
			   / \
			  B   C
			 / \   \
			D   E - F
		*/
		order := make([]string, 0)
		onVisit := func(node string) {
			order = append(order, node)
		}
		err := g.DFS("A", onVisit)
		assertx.Nil(t, err)
		assertx.Equal(t, order, []string{"A", "B", "D", "E", "F", "C"})
	})
}

func TestGraph_Dijkstra(t *testing.T) {
	t.Run("dijkstra on empty graph yields error", func(t *testing.T) {
		g := New[string]()
		path, distance, err := g.Dijkstra("start", "end")
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
	g := New[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 4)
	g.AddEdge("B", "C", 1)
	g.AddEdge("B", "D", 5)
	g.AddEdge("C", "D", 2)
	/*
	   A --1-- B
	   |     / |
	   |    /  |
	   4   1   5
	   |  /    |
	   | /     |
	   C --2-- D
	*/
	t.Run("one node path", func(t *testing.T) {
		path, distance, err := g.Dijkstra("A", "B")
		assertx.Nil(t, err)
		assertx.Equal(t, path, []string{"A", "B"})
		assertx.Equal(t, distance, 1)
	})
	t.Run("end node not present in graph", func(t *testing.T) {
		path, distance, err := g.Dijkstra("A", "Z")
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
	t.Run("shortest path from A to D is correct with weight 4", func(t *testing.T) {
		// A -> B -> C -> D weight: 4 is shorter than A -> B -> D weight: 6
		path, distance, err := g.Dijkstra("A", "D")
		assertx.Nil(t, err)
		assertx.Equal(t, path, []string{"A", "B", "C", "D"})
		assertx.Equal(t, distance, 4)
	})
	t.Run("returns error for an unreachable node in a disconnected graph", func(t *testing.T) {
		g.AddNode("Z")
		path, distance, err := g.Dijkstra("A", "Z")
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
}

type coord struct {
	x int
	y int
}

func manhattanDistance(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func euclideanDistance(a, b coord) int {
	dx := b.x - a.x
	dy := b.y - a.y
	return dx*dx + dy*dy
}

func chebyshevDistance(a, b coord) int {
	dx := abs(b.x - a.x)
	dy := abs(b.y - a.y)
	return max(dx, dy)
}

func TestGraph_AStar(t *testing.T) {
	origin := coord{x: 0, y: 0}
	t.Run("astar on empty graph yields error", func(t *testing.T) {
		g := New[coord]()
		path, distance, err := g.AStar(coord{x: 0, y: 0}, coord{x: 1, y: 1}, manhattanDistance)
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
	t.Run("end node not present in graph", func(t *testing.T) {
		g := New[coord]()
		g.AddNode(origin)
		path, distance, err := g.AStar(origin, coord{x: 1, y: 1}, manhattanDistance)
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
	t.Run("disconnected start and end yields no path", func(t *testing.T) {
		g := New[coord]()
		g.AddNode(origin)
		g.AddNode(coord{x: 2, y: 2})
		path, distance, err := g.AStar(origin, coord{x: 2, y: 2}, manhattanDistance)
		assertx.NotNil(t, err)
		assertx.Nil(t, path)
		assertx.Equal(t, distance, 0)
	})
	t.Run("valid path using manhattan distance", func(t *testing.T) {
		g := New[coord]()
		coordB := coord{x: 1, y: 0}
		coordC := coord{x: 1, y: 1}
		g.AddEdge(origin, coordB, 1)
		g.AddEdge(coordB, coordC, 1)
		g.AddEdge(origin, coordC, 10)
		path, weight, err := g.AStar(origin, coordC, manhattanDistance)
		assertx.Nil(t, err)
		assertx.Equal(t, path, []coord{origin, coordB, coordC})
		assertx.Equal(t, weight, 2)
	})
	t.Run("valid path using euclidean distance for `valley` graph", func(t *testing.T) {
		g := New[coord]()
		a := coord{x: 0, y: 2}
		b := coord{x: 1, y: 2}
		c := coord{x: 2, y: 2}
		d := coord{x: 0, y: 0}
		e := coord{x: 1, y: 0}
		f := coord{x: 2, y: 0}
		// high weight but direct
		g.AddEdge(a, b, 100)
		g.AddEdge(b, c, 1)
		// low weight but long path
		g.AddEdge(a, d, 1)
		g.AddEdge(d, e, 1)
		g.AddEdge(e, f, 1)
		g.AddEdge(f, c, 1)
		path, weight, err := g.AStar(a, c, euclideanDistance)
		assertx.Nil(t, err)
		assertx.Equal(t, weight, 4)
		assertx.Equal(t, path, []coord{a, d, e, f, c})
	})
	t.Run("valid diagonal path using chebyshev distance", func(t *testing.T) {
		g := New[coord]()
		a := coord{x: 0, y: 2}
		b := coord{x: 0, y: 0}
		c := coord{x: 2, y: 2}
		d := coord{x: 2, y: 0}
		g.AddEdge(a, b, 1)
		g.AddEdge(b, d, 1)
		g.AddEdge(a, c, 10)
		g.AddEdge(c, d, 1)
		g.AddEdge(a, d, 1)
		path, weight, err := g.AStar(a, d, chebyshevDistance)
		assertx.Nil(t, err)
		assertx.Equal(t, weight, 1)
		assertx.Equal(t, path, []coord{a, d})
	})
}

func TestGraph_String(t *testing.T) {
	t.Run("string is empty for empty graph", func(t *testing.T) {
		g := New[string]()
		assertx.Equal(t, g.String(), "")
	})
	t.Run("graph with one node", func(t *testing.T) {
		g := New[string]()
		g.AddNode("London")
		assertx.Equal(t, g.String(), "London")
	})
	t.Run("graph with multiple nodes and edges", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 1)
		g.AddEdge("C", "A", 2)
		g.AddEdge("D", "A", 3)
		expected := "A -> B (1)" + "\n" + "B" + "\n" + "C -> A (2)" + "\n" + "D -> A (3)"
		assertx.Equal(t, g.String(), expected)
	})
	t.Run("graph with multiple outgoing edges is formatted correctly", func(t *testing.T) {
		g := New[string]()
		g.AddEdge("A", "B", 1)
		g.AddEdge("A", "C", 2)
		g.AddEdge("C", "A", 3)
		expected := "A -> B (1), C (2)" + "\n" + "B" + "\n" + "C -> A (3)"
		assertx.Equal(t, g.String(), expected)
	})
}

func createGridGraph(size int) (*Graph[coord], coord, coord) {
	g := New[coord]()
	for y := range size {
		for x := range size {
			c := coord{x: x, y: y}
			if x+1 < size {
				g.AddEdge(c, coord{x: x + 1, y: y}, rand.IntN(10)+1)
			}
			if y+1 < size {
				g.AddEdge(c, coord{x: x, y: y + 1}, rand.IntN(10)+1)
			}
		}
	}
	start := coord{x: 0, y: 0}
	end := coord{x: size - 1, y: size - 1}
	return g, start, end
}

func BenchmarkBFS_100(b *testing.B) {
	size := 100
	g, start, _ := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.BFS(start, func(node coord) {})
	}
}

func BenchmarkBFS_1_000(b *testing.B) {
	size := 1_000
	g, start, _ := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.BFS(start, func(node coord) {})
	}
}

func BenchmarkDFS_100(b *testing.B) {
	size := 100
	g, start, _ := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.DFS(start, func(node coord) {})
	}
}

func BenchmarkDFS_1_000(b *testing.B) {
	size := 1_000
	g, start, _ := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.DFS(start, func(node coord) {})
	}
}

func BenchmarkDijkstra_100(b *testing.B) {
	size := 100
	g, start, end := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.Dijkstra(start, end)
	}
}

func BenchmarkDijkstra_1_000(b *testing.B) {
	size := 1_000
	g, start, end := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.Dijkstra(start, end)
	}
}

func BenchmarkAStar_100(b *testing.B) {
	size := 100
	g, start, end := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.AStar(start, end, manhattanDistance)
	}
}

func BenchmarkAStar_1_000(b *testing.B) {
	size := 1_000
	g, start, end := createGridGraph(size)
	b.ResetTimer()
	for b.Loop() {
		g.AStar(start, end, manhattanDistance)
	}
}
