package filetree

import "sort"

type NodeType string

const (
	FILE NodeType = "FILE"
	DIR  NodeType = "DIR"
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
