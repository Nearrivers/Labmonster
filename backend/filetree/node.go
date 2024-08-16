package filetree

import (
	"time"
)

type NodeType string
type FileType string

const (
	FILE NodeType = "FILE"
	DIR  NodeType = "DIR"
)

const (
	GRAPH FileType = "GRAPH"
	SHEET FileType = "SHEET"
	VIDEO FileType = "VIDEO"
	IMAGE FileType = "IMAGE"
)

// A node is the in-memory representation of a file or a directory on the user's machine
type Node struct {
	Name      string    `json:"name"`
	Type      NodeType  `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	Extension string    `json:"extension"`
	FileType  FileType  `json:"fileType"`
}

func NewNode(name, extension string, nodeType NodeType) Node {
	return Node{
		Name:      name,
		Type:      nodeType,
		Extension: extension,
		FileType:  GRAPH,
		UpdatedAt: time.Now(),
	}
}
