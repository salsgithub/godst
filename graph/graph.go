package graph

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/salsgithub/godst/heap"
)

type Edge[T comparable] struct {
	Link   T
	Weight int
}

type Graph[T comparable] struct {
	adjacency map[T][]Edge[T]
}

func New[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjacency: make(map[T][]Edge[T]),
	}
}

func (g *Graph[T]) AddNode(value T) {
	if _, ok := g.adjacency[value]; ok {
		return
	}
	g.adjacency[value] = []Edge[T]{}
}

func (g *Graph[T]) AddEdge(from, to T, weight int) {
	g.AddNode(from)
	g.AddNode(to)
	edge := Edge[T]{
		Link:   to,
		Weight: weight,
	}
	g.adjacency[from] = append(g.adjacency[from], edge)
}

func (g *Graph[T]) DeleteNode(value T) {
	if _, ok := g.adjacency[value]; !ok {
		return
	}
	delete(g.adjacency, value)
	for node, edges := range g.adjacency {
		newEdges := []Edge[T]{}
		for _, edge := range edges {
			if edge.Link != value {
				newEdges = append(newEdges, edge)
			}
		}
		g.adjacency[node] = newEdges
	}
}

func (g *Graph[T]) Neighbours(value T) ([]Edge[T], bool) {
	neighbours, ok := g.adjacency[value]
	return neighbours, ok
}

type visitedState byte

const (
	unvisited visitedState = iota
	visiting
	visited
)

func (g *Graph[T]) HasCycle() bool {
	states := make(map[T]visitedState)
	for node := range g.adjacency {
		states[node] = unvisited
	}
	for node := range g.adjacency {
		if states[node] == unvisited {
			if checkCycle(g, node, states) {
				return true
			}
		}
	}
	return false
}

func checkCycle[T comparable](g *Graph[T], node T, states map[T]visitedState) bool {
	states[node] = visiting
	neighbours, _ := g.Neighbours(node)
	for _, neighbour := range neighbours {
		link := neighbour.Link
		switch states[link] {
		case visiting:
			return true
		case unvisited:
			if checkCycle(g, link, states) {
				return true
			}
		}
	}
	states[node] = visited
	return false
}

func (g *Graph[T]) TopologicalSort() ([]T, error) {
	if g.HasCycle() {
		return nil, errors.New("graph contains cycle for topological sorting")
	}
	result := make([]T, 0)
	visited := make(map[T]bool)
	var dfs func(T)
	dfs = func(node T) {
		visited[node] = true
		neighbours, _ := g.Neighbours(node)
		for _, neighbour := range neighbours {
			if !visited[neighbour.Link] {
				dfs(neighbour.Link)
			}
		}
		result = append([]T{node}, result...)
	}
	for _, node := range g.Nodes() {
		if !visited[node] {
			dfs(node)
		}
	}
	return result, nil
}

func (g *Graph[T]) BFS(start T, onVisit func(node T)) error {
	if _, ok := g.adjacency[start]; !ok {
		return fmt.Errorf("start %v not found in graph", start)
	}
	queue := []T{start}
	visited := map[T]bool{
		start: true,
	}
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		onVisit(next)
		neighbours, _ := g.Neighbours(next)
		for _, neighbour := range neighbours {
			link := neighbour.Link
			if !visited[link] {
				visited[link] = true
				queue = append(queue, link)
			}
		}
	}
	return nil
}

func (g *Graph[T]) DFS(start T, onVisit func(node T)) error {
	if _, ok := g.adjacency[start]; !ok {
		return fmt.Errorf("start %v not found in graph", start)
	}
	queue := []T{start}
	visited := make(map[T]bool)
	for len(queue) > 0 {
		last := len(queue) - 1
		node := queue[last]
		queue = queue[:last]
		if !visited[node] {
			visited[node] = true
			onVisit(node)
			neighbours, _ := g.Neighbours(node)
			for i := len(neighbours) - 1; i >= 0; i-- {
				link := neighbours[i].Link
				if !visited[link] {
					queue = append(queue, link)
				}
			}
		}
	}
	return nil
}

type priorityNode[T comparable] struct {
	node     T
	priority int
}

func (g *Graph[T]) Dijkstra(start, end T) ([]T, int, error) {
	if _, ok := g.adjacency[start]; !ok {
		return nil, 0, fmt.Errorf("start node %v not found", start)
	}
	if _, ok := g.adjacency[end]; !ok {
		return nil, 0, fmt.Errorf("end node %v not found", end)
	}
	distances := make(map[T]int)
	route := make(map[T]T)
	queue := heap.New(func(a, b priorityNode[T]) bool {
		return a.priority < b.priority
	})
	for node := range g.adjacency {
		distances[node] = math.MaxInt
	}
	distances[start] = 0
	queue.Push(priorityNode[T]{node: start, priority: 0})
	for !queue.IsEmpty() {
		pop, _ := queue.Pop()
		node := pop.node
		if pop.priority > distances[node] {
			continue
		}
		if node == end {
			break
		}
		neighbours, _ := g.Neighbours(node)
		for _, neighbour := range neighbours {
			link := neighbour.Link
			weight := neighbour.Weight
			travelDistance := distances[node] + weight
			if travelDistance < distances[link] {
				distances[link] = travelDistance
				route[link] = node
				queue.Push(priorityNode[T]{node: link, priority: travelDistance})
			}
		}
	}
	if distances[end] == math.MaxInt {
		return nil, 0, fmt.Errorf("path from %v to %v not found", start, end)
	}
	path := []T{}
	n := end
	for {
		path = append([]T{n}, path...)
		step, ok := route[n]
		if !ok {
			break
		}
		n = step
	}
	return path, distances[end], nil
}

func (g *Graph[T]) AStar(start, end T, heuristic func(a, b T) int) ([]T, int, error) {
	if _, ok := g.adjacency[start]; !ok {
		return nil, 0, fmt.Errorf("start node %v not found", start)
	}
	if _, ok := g.adjacency[end]; !ok {
		return nil, 0, fmt.Errorf("end node %v not found", end)
	}
	scoreG := make(map[T]int)
	scoreF := make(map[T]int)
	for node := range g.adjacency {
		scoreG[node] = math.MaxInt
	}
	scoreG[start] = 0
	scoreF[start] = heuristic(start, end)
	route := make(map[T]T)
	queue := heap.New(func(a, b priorityNode[T]) bool {
		return a.priority < b.priority
	})
	queue.Push(priorityNode[T]{node: start, priority: scoreF[start]})
	for !queue.IsEmpty() {
		pop, _ := queue.Pop()
		node := pop.node
		if node == end {
			path := []T{}
			current := end
			for {
				path = append([]T{current}, path...)
				step, ok := route[current]
				if !ok {
					break
				}
				current = step
			}
			return path, scoreG[end], nil
		}
		neighbours, _ := g.Neighbours(node)
		for _, neighbour := range neighbours {
			link := neighbour.Link
			scoreTentative := scoreG[node] + neighbour.Weight
			if scoreTentative < scoreG[link] {
				route[link] = node
				scoreG[link] = scoreTentative
				scoreF[link] = scoreG[link] + heuristic(link, end)
				queue.Push(priorityNode[T]{node: link, priority: scoreF[link]})
			}
		}
	}
	return nil, 0, fmt.Errorf("path from %v to %v not found", start, end)
}

func (g *Graph[T]) Nodes() []T {
	nodes := make([]T, 0, len(g.adjacency))
	for node := range g.adjacency {
		nodes = append(nodes, node)
	}
	a := any(nodes)
	switch t := a.(type) {
	case []string:
		slices.Sort(t)
	case []int:
		slices.Sort(t)
	case []float64:
		slices.Sort(t)
	}
	return nodes
}

func (g *Graph[T]) Len() int {
	return len(g.adjacency)
}

func (g *Graph[T]) String() string {
	if g.Len() == 0 {
		return ""
	}
	builder := strings.Builder{}
	nodes := g.Nodes()
	for i, node := range nodes {
		builder.WriteString(fmt.Sprintf("%v", node))
		neighbours := g.adjacency[node]
		if len(neighbours) > 0 {
			builder.WriteString(" -> ")
			for j, edge := range neighbours {
				builder.WriteString(fmt.Sprintf("%v (%d)", edge.Link, edge.Weight))
				if j < len(neighbours)-1 {
					builder.WriteString(", ")
				}
			}
		}
		if i < len(nodes)-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
