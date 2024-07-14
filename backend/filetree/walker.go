package filetree

import (
	"fmt"
	"sort"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// Implement the sort interface
type NodeSequences []Node

func (ns NodeSequences) Len() int {
	return len(ns)
}

func (ns NodeSequences) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

// Nodes are sorted in alphabetical order
func (ns NodeSequences) Less(i, j int) bool {
	return ns[i].Name < ns[j].Name
}

func Walk(root *Node, ch chan Node) {
	defer close(ch)
	if root != nil {
		ch <- Node{
			Name: root.Name,
			Type: root.Type,
		}
		for _, fileNode := range root.Files {
			walkRecursively(fileNode, ch)
		}
	}
}

func walkRecursively(node *Node, ch chan Node) {
	if node != nil {
		ch <- Node{
			Name: node.Name,
			Type: node.Type,
		}
		for _, fileNode := range node.Files {
			walkRecursively(fileNode, ch)
		}
	}
}

// Print the tree for debugging purposes
func PrintTree(fileTree Node) {
	ch1 := make(chan Node)

	go Walk(&fileTree, ch1)

	for n1 := range ch1 {
		logger := logger.NewDefaultLogger()
		logger.Debug(fmt.Sprintf("Name: %s, Type: %s", n1.Name, n1.Type))
	}
}

// Used in tests (for now). Compare 2 trees and tells whether or not they are identicals
func Same(fileTree, otherFileTree *Node) bool {
	ch1 := make(chan Node)
	ch2 := make(chan Node)

	namesInR1 := NodeSequences{}
	namesInR2 := NodeSequences{}

	go Walk(fileTree, ch1)
	go Walk(otherFileTree, ch2)

	for n1 := range ch1 {
		namesInR1 = append(namesInR1, n1)
	}

	for n2 := range ch2 {
		namesInR2 = append(namesInR2, n2)
	}

	sort.Sort(namesInR1)
	sort.Sort(namesInR2)

	if len(namesInR1) != len(namesInR2) {
		return false
	}

	for i, n := range namesInR1 {
		if n.Name != namesInR2[i].Name {
			return false
		}
	}

	return true
}
