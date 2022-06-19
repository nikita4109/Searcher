package aho_corasick

import (
	"io"
)

type AhoCorasick struct {
	root *Vertex
}

func New() AhoCorasick {
	root := NewVertex(nil, -1, nil)
	root.root = root

	return AhoCorasick{root: root}
}

func (aho *AhoCorasick) Add(s *string) {
	vertex := aho.root

	for c := range *s {
		if _, ok := vertex.direct[c]; !ok {
			vertex.direct[c] = NewVertex(aho.root, c, vertex)
		}

		vertex = vertex.direct[c]
	}

	if vertex != aho.root {
		vertex.termination = true
	}
}

func (aho *AhoCorasick) Contain(reader io.RuneReader) bool {
	vertex := aho.root

	for {
		character, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}

		vertex = vertex.crossByCharacter(int(character))
		if vertex.termination {
			return true
		}
	}

	return false
}
