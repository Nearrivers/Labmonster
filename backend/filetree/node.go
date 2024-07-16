package filetree

import (
	"errors"
	"sort"
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
	// A map keeping the indexes of the Files array. Used in order to find nodes fast
	NameToIndex map[string]int
}

func SortNodes(files []*Node) {
	sort.SliceStable(files, func(i, j int) bool {
		if files[i].Type != files[j].Type {
			return files[i].Type < files[j].Type
		}

		return files[i].Name < files[j].Name
	})
}

func InsertNode(isDir bool, node *Node, name string) *Node {
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

	node.Files = append(node.Files, &newNode)

	return &newNode
}

func removeIndex(n []*Node, index int) ([]*Node, error) {
	if index >= len(n) {
		return nil, ErrIndexOutOfBounds
	}

	ret := make([]*Node, 0)
	ret = append(ret, n[:index]...)
	return append(ret, n[index+1:]...), nil
}
