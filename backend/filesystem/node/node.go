package node

import (
	"io/fs"
	"path/filepath"
	"sort"
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

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	iNode, jNode := n[i], n[j]

	if iNode.Type == jNode.Type {
		return iNode.Name < jNode.Name
	}

	if iNode.Type == DIR {
		return true
	}

	return false
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
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
	dirNames := make([]*Node, 0)
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

	sort.Sort(Nodes(dirNames))
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
