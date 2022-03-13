package ds

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
