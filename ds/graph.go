package ds

type Vertex struct {
	Name             string
	AdjacentVertices map[string]*Vertex
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:             name,
		AdjacentVertices: make(map[string]*Vertex),
	}
}

func (v *Vertex) addAdjacentVertex(adjV *Vertex) {
	if _, ok := v.AdjacentVertices[adjV.Name]; !ok {
		v.AdjacentVertices[adjV.Name] = adjV
	}

	return
}
