package node

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
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

type Nodes []*Node

func (n Nodes) String() string {
	s := "["
	for i, node := range n {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", node)
	}

	return s + "]"
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

// Takes an array of fs.DirEntry to create an array of type *Node and returns it.
// This function ignore any element that starts with a dot (.git or .labmonster for example) as
// it contains config files that are not relevant to the user
func CreateNodesFromDirEntries(entries []fs.DirEntry) (Nodes, error) {
	dirNames := make(Nodes, 0)
	for _, entry := range entries {
		ext := filepath.Ext(entry.Name())
		if entry.Name() == ext {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		newNode := Node{
			Name:      strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())),
			Extension: filepath.Ext(entry.Name()),
			UpdatedAt: info.ModTime(),
		}

		if entry.IsDir() {
			newNode.Type = DIR
		} else {
			newNode.FileType = DetectFileType(ext)
			newNode.Type = FILE
		}

		dirNames = append(dirNames, &newNode)
	}

	slices.SortStableFunc(dirNames, func(iNode, jNode *Node) int {
		if iNode.Type == jNode.Type {
			return strings.Compare(strings.ToLower(iNode.Name), strings.ToLower(jNode.Name))
		}

		if iNode.Type == DIR {
			return -1
		}

		return 1
	})

	return dirNames, nil
}

// Given an extension, it wil return the corresponding FileType
func DetectFileType(extension string) FileType {
	switch extension {
	case ".png", ".jpeg", ".gif", ".webp":
		return IMAGE
	case ".json":
		return GRAPH
	case ".mp4", ".mpeg":
		return VIDEO
	default:
		return UNSUPPORTED
	}
}
