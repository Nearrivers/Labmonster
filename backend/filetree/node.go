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

type Node struct {
	Name  string   `json:"name"`
	Type  NodeType `json:"type"`
	Files []*Node  `json:"files"`
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

func RemoveNode(node *Node, name string) (*Node, error) {
	_, index, err := searchFile(name, node.Files)
	if err != nil {
		return nil, err
	}

	newFiles, err := removeIndex(node.Files, index)
	if err != nil {
		return nil, err
	}

	node.Files = newFiles
	return node, nil
}

func removeIndex(n []*Node, index int) ([]*Node, error) {
	if index >= len(n) {
		return nil, ErrIndexOutOfBounds
	}

	ret := make([]*Node, 0)
	ret = append(ret, n[:index]...)
	return append(ret, n[index+1:]...), nil
}
