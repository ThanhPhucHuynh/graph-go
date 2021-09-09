package main

import "fmt"

type Vertex struct {
	Key int

	Vertices map[int]*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

// Graph
type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}
func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func DFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	visited := map[int]bool{}

	if startVertex == nil {
		return
	}
	visited[startVertex.Key] = true
	visitCb(startVertex.Key)
	for _, v := range startVertex.Vertices {
		if visited[v.Key] {
			continue
		}
		DFS(g, v, visitCb)
	}

	// fmt.Println(visited)
}

type node struct {
	v    *Vertex
	next *node
}

type queue struct {
	head *node
	tail *node
}

func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() *Vertex {
	n := q.head
	// return nil, if head is empty
	if n == nil {
		return nil
	}

	q.head = q.head.next

	// 	// if there wasn't any next node, that
	// 	// means the queue is empty, and the tail
	// 	// should be set to nil
	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

func BFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}

	currentVertex := startVertex

	for {
		visitCb(currentVertex.Key)
		visitedVertices[currentVertex.Key] = true

		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		currentVertex = vertexQueue.dequeue()
		if currentVertex == nil {
			break
		}
	}

}

func main() {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	BFS(g, g.Vertices[1], cb)

	// add assertions here
	fmt.Println(visitedOrder)

	fmt.Println("Dijkstra")
	// Example
	graph := newGraph()
	graph.addEdge("S", "B", 4)
	graph.addEdge("S", "C", 2)
	graph.addEdge("B", "C", 1)
	graph.addEdge("B", "D", 5)
	graph.addEdge("C", "D", 8)
	graph.addEdge("C", "E", 10)
	graph.addEdge("D", "E", 2)
	graph.addEdge("D", "T", 6)
	graph.addEdge("E", "T", 2)
	fmt.Println(graph.getPath("S", "T"))

}
