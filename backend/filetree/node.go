package filetree

import (
	"errors"
	"sort"
	"strings"
)

type NodeType string

const (
	FILE NodeType = "FILE"
	DIR  NodeType = "DIR"
)

var (
	ErrIndexOutOfBounds = errors.New("l'index donné est trop grand")
)

// A node is the in-memory representation of a file or a directory on the user's machine
type Node struct {
	Name  string   `json:"name"`
	Type  NodeType `json:"type"`
	Files []*Node  `json:"files"`
}

func (n *Node) SetName(newName string) {
	if !strings.HasSuffix(newName, ".json") {
		n.Name = newName + ".json"
		return
	}

	n.Name = newName
}

func (n *Node) SortNodes() {
	sort.SliceStable(n.Files, func(i, j int) bool {
		n1, n2 := n.Files[i], n.Files[j]

		if n1.Type != n2.Type {
			return n1.Type < n2.Type
		}

		return n1.Name < n2.Name
	})
}

func (n *Node) InsertNode(isDir bool, name string) *Node {
	var nodetype NodeType

	if isDir {
		nodetype = DIR
	} else {
		nodetype = FILE
	}

	newNode := Node{
		Name:  name,
		Type:  nodetype,
		Files: []*Node{},
	}

	n.Files = append(n.Files, &newNode)
	n.SortNodes()
	return &newNode
}

func (n *Node) removeIndex(index int) error {
	if index >= len(n.Files) {
		return ErrIndexOutOfBounds
	}

	ret := make([]*Node, 0)
	ret = append(ret, n.Files[:index]...)
	n.Files = append(ret, n.Files[index+1:]...)

	return nil
}
