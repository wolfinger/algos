package ds

import "fmt"

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Name             string
	AdjacentVertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(v *Vertex) {
	if _, ok := g.Vertices[v.Name]; !ok {
		g.Vertices[v.Name] = v
	}

	return
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:             name,
		AdjacentVertices: make(map[string]*Vertex),
	}
}

func (v *Vertex) AddAdjacentVertex(adjV *Vertex) {
	if _, ok := v.AdjacentVertices[adjV.Name]; !ok {
		v.AdjacentVertices[adjV.Name] = adjV
	}

	return
}

func (g *Graph) DFSTraverse(v *Vertex, visited map[string]bool) {
	visited[v.Name] = true

	fmt.Println("visited: " + v.Name)

	for key, adjV := range v.AdjacentVertices {
		if _, ok := visited[key]; !ok {
			g.DFSTraverse(adjV, visited)
		}
	}

	return
}

func (g *Graph) BFSTraverese(v *Vertex) {
	var vq VertexQueue
	visited := make(map[string]bool)

	visited[v.Name] = true
	vq.Enqueue(*v)

	for !vq.IsEmpty() {
		currV, _ := vq.Dequeue()

		fmt.Println("visited: ", currV.Name)

		for key, adjV := range currV.AdjacentVertices {
			if _, ok := visited[key]; !ok {
				visited[key] = true
				vq.Enqueue(*adjV)
			}
		}
	}

	return
}

type VertexQueue struct {
	queue []Vertex
}

func (s *VertexQueue) Enqueue(item Vertex) {
	// add to queue
	s.queue = append(s.queue, item)
}

func (s *VertexQueue) Dequeue() (Vertex, bool) {
	// remove from queue
	if len(s.queue) > 0 {
		item := s.queue[0]
		s.queue = s.queue[1:len(s.queue)]

		return item, false
	} else {
		return Vertex{}, true
	}
}

func (s *VertexQueue) Read() (Vertex, bool) {
	// read from queue
	if len(s.queue) == 0 {
		return Vertex{}, true
	} else {
		return s.queue[0], false
	}
}

func (s *VertexQueue) IsEmpty() bool {
	if len(s.queue) == 0 {
		return true
	} else {
		return false
	}
}
