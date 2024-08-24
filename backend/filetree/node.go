package filetree

import (
	"time"
)

type DataType string
type FileType string

const (
	FILE DataType = "FILE"
	DIR  DataType = "DIR"
)

var DTypes = []struct {
	Value  DataType
	TSName string
}{
	{FILE, "FILE"},
	{DIR, "DIR"},
}

const (
	GRAPH       FileType = "GRAPH"
	SHEET       FileType = "SHEET"
	VIDEO       FileType = "VIDEO"
	IMAGE       FileType = "IMAGE"
	UNSUPPORTED FileType = "UNSUPPORTED"
)

var FTypes = []struct {
	Value  FileType
	TSName string
}{
	{GRAPH, "GRAPH"},
	{SHEET, "SHEET"},
	{VIDEO, "VIDEO"},
	{IMAGE, "IMAGE"},
	{UNSUPPORTED, "UNSUPPORTED"},
}

// A node is the in-memory representation of a file or a directory on the user's machine
type Node struct {
	Name      string    `json:"name"`
	Type      DataType  `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	Extension string    `json:"extension"`
	FileType  FileType  `json:"fileType"`
}

func NewNode(name, extension string, nodeType DataType) Node {
	return Node{
		Name:      name,
		Type:      nodeType,
		Extension: extension,
		FileType:  GRAPH,
		UpdatedAt: time.Now(),
	}
}
