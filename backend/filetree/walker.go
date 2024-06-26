package filetree

import (
	"fmt"
	"sort"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// Implemente l'interface sort
type NodeSequences []Node

func (ns NodeSequences) Len() int {
	return len(ns)
}

func (ns NodeSequences) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

// Les noeuds sont triés dans l'ordre alphabétique
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

// Permet de log l'arbre dans la console pour débugger
func PrintTree(fileTree Node) {
	ch1 := make(chan Node)

	go Walk(&fileTree, ch1)

	for n1 := range ch1 {
		logger := logger.NewDefaultLogger()
		logger.Debug(fmt.Sprintf("Nom: %s, Type: %s", n1.Name, n1.Type))
	}
}

// Fonction utilisée pour les tests. Permet de comparer 2 arbres et de retourner si oui ou non ils sont identiques
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
