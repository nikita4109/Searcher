package aho_corasick

type Vertex struct {
	direct, next map[int]*Vertex
	link, parent *Vertex

	parentCharacter int
	termination     bool

	root *Vertex
}

func NewVertex(root *Vertex, pCh int, parent *Vertex) *Vertex {
	return &Vertex{
		direct:          make(map[int]*Vertex),
		next:            make(map[int]*Vertex),
		root:            root,
		parentCharacter: pCh,
		parent:          parent,
	}
}

func (vertex *Vertex) getLink() *Vertex {
	if vertex.link == nil {
		if vertex == vertex.root || vertex.parent == vertex.root {
			vertex.link = vertex.root
		} else {
			vertex.link = (vertex.parent.getLink()).crossByCharacter(vertex.parentCharacter)
		}
	}

	return vertex.link
}

func (vertex *Vertex) crossByCharacter(character int) *Vertex {
	if _, ok := vertex.next[character]; !ok {
		if _, ok := vertex.direct[character]; ok {
			vertex.next[character] = vertex.direct[character]
		} else if vertex == vertex.root {
			vertex.next[character] = vertex.root
		} else {
			vertex.next[character] = vertex.getLink().crossByCharacter(character)
		}
	}

	return vertex.next[character]
}
