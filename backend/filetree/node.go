package filetree

import (
	"errors"
	"sort"
	"strings"
	"time"
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
	Name      string    `json:"name"`
	Type      NodeType  `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Extension string    `json:"extension"`
	Files     []*Node   `json:"files"`
}

func NewNode(name, extension string, nodeType NodeType) Node {
	return Node{
		Name:      name,
		Type:      nodeType,
		Extension: extension,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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
