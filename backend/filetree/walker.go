package filetree

import (
	"fmt"
	"sort"
)

type NodeSequence struct {
	Name string
	Type NodeType
}

type NodeSequences []NodeSequence

func (ns NodeSequences) Len() int {
	return len(ns)
}

func (ns NodeSequences) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns NodeSequences) Less(i, j int) bool {
	return ns[i].Name < ns[j].Name
}

func Walk(root *Node, ch chan NodeSequence) {
	defer close(ch)
	if root != nil {
		ch <- NodeSequence{
			Name: root.Name,
			Type: root.Type,
		}
		for _, fileNode := range root.Files {
			walkRecursively(fileNode, ch)
		}
	}
}

func walkRecursively(node *Node, ch chan NodeSequence) {
	if node != nil {
		ch <- NodeSequence{
			Name: node.Name,
			Type: node.Type,
		}
		for _, fileNode := range node.Files {
			walkRecursively(fileNode, ch)
		}
	}
}

// Permet de log l'arbre dans la console pour débugger
func (ft *FileTreeExplorer) PrintTree() {
	ch1 := make(chan NodeSequence)

	go Walk(&ft.FileTree, ch1)

	for n1 := range ch1 {
		ft.Logger.Debug(fmt.Sprintf("Nom: %s, Type: %s", n1.Name, n1.Type))
	}
}

// Fonction utilisée pour les tests. Permet de comparer 2 arbres et de retourner si oui ou non ils sont identiques
func (ft *FileTreeExplorer) Same(otherFileTree *Node) bool {
	ch1 := make(chan NodeSequence)
	ch2 := make(chan NodeSequence)

	namesInR1 := NodeSequences{}
	namesInR2 := NodeSequences{}

	go Walk(&ft.FileTree, ch1)
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
		if n != namesInR2[i] {
			return false
		}
	}

	return true
}
