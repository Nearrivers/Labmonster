package filetree

import (
	"errors"
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
