package filetree

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"flow-poc/backend/config"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

var (
	ErrPathMissing = errors.New("the path must be specified")
)

type FileTreeExplorer struct {
	Logger   logger.Logger     `json:"logger"`
	Cfg      *config.AppConfig `json:"cfg"`
	FileTree Node              `json:"file_tree"`
}

func NewFileTree(cfg *config.AppConfig) *FileTreeExplorer {
	return &FileTreeExplorer{
		Logger: logger.NewDefaultLogger(),
		Cfg:    cfg,
	}
}

func (ft *FileTreeExplorer) GetLabPath() string {
	return ft.Cfg.ConfigFile.LabPath
}

// Given a path to a directory starting from the lab root, this function will read
// its content using the os.ReadDir method, transforms those entries into Nodes
// and return them
func (ft *FileTreeExplorer) GetSubDirAndFiles(pathFromLabRoot string) ([]*Node, error) {
	dirPath := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	nodes := ft.createNodesFromDirEntries(entries)

	// if len(nodes) > 0 {
	// 	err = ft.SetNodeFiles(pathFromLabRoot, nodes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return nodes, nil
}

// Given a path to a directory starting from the lab root and an array of Nodes,
// this function will find the node corresponding to the path and set its Files
// attribute with the value of the nodes argument
func (ft *FileTreeExplorer) SetNodeFiles(pathFromLabRoot string, nodes []*Node) error {
	if pathFromLabRoot == "" {
		ft.FileTree.Files = nodes
		return nil
	}

	n, _, err := ft.FindNodeWithPath(pathFromLabRoot)
	if err != nil {
		return err
	}

	n.Files = nodes
	return nil
}

// Given a path to a node starting from the lab root, this function will, first, find its
// parent node to then delete it
func (ft *FileTreeExplorer) RemoveNode(pathFromLabRoot string) error {
	if pathFromLabRoot == "" {
		return ErrPathMissing
	}

	separator := string(filepath.Separator)
	path := strings.Split(pathFromLabRoot, separator)

	if len(path) == 1 {
		err := ft.findAndDeleteNode(pathFromLabRoot, ft.FileTree.Files)
		if err != nil {
			return nil
		}

		return nil
	}

	pathToParentFolder := strings.Join(path[:len(path)-2], separator)
	fileName := path[len(path)-1]

	n, _, err := ft.FindNodeWithPath(pathToParentFolder)
	if err != nil {
		return err
	}

	err = ft.findAndDeleteNode(fileName, n.Files)
	if err != nil {
		return err
	}

	return nil
}

func (ft *FileTreeExplorer) RenameNode(pathFromLabRoot, oldName, newName string) error {
	addJsonSuffix(&oldName)
	addJsonSuffix(&newName)

	if pathFromLabRoot == "" {
		err := ft.findAndRenameNode(oldName, newName, ft.FileTree.Files)
		if err != nil {
			return err
		}

		return nil
	}

	f, _, err := ft.FindNodeWithPath(pathFromLabRoot)
	if err != nil {
		return err
	}

	err = ft.findAndRenameNode(oldName, newName, f.Files)
	if err != nil {
		return err
	}

	return nil
}

// Given a node name and a node array that should contains it, deletes the node and returns
// the node array without it
func (ft *FileTreeExplorer) findAndDeleteNode(nodeName string, files []*Node) error {
	n, i, err := searchFileOrDir(nodeName, files)
	if err != nil {
		return err
	}

	err = n.removeIndex(i)
	if err != nil {
		return err
	}

	return nil
}

func (ft *FileTreeExplorer) findAndRenameNode(oldName, newName string, f []*Node) error {
	n, _, err := searchFileOrDir(oldName, f)
	if err != nil {
		return err
	}

	n.SetName(newName)
	return nil
}

// Given a path starting from the lab root, return the corresponding node
// by splitting the argument string by filepath.Separator and binary searching
// each node by its name
func (ft *FileTreeExplorer) FindNodeWithPath(pathFromLabRoot string) (*Node, int, error) {
	separator := string(filepath.Separator)
	path := strings.Split(pathFromLabRoot, separator)
	n := &ft.FileTree
	index := 0

	for _, dir := range path {
		n.SortNodes()
		node, i, err := searchFileOrDir(dir, n.Files)
		if err != nil {
			return nil, -1, err
		}

		n = node
		index = i
	}

	return n, index, nil
}

// Takes an array of fs.DirEntry to create an array of type *Node and returns it
// This function ignore any .labmonster directory as it might contain config files that are not relevant to the user
func (ft *FileTreeExplorer) createNodesFromDirEntries(entries []fs.DirEntry) []*Node {
	dirNames := make([]*Node, 0)
	for _, entry := range entries {
		if entry.Name() == ".labmonster" && entry.IsDir() {
			continue
		}

		newNode := Node{
			Name:  entry.Name(),
			Files: make([]*Node, 0),
		}

		if entry.IsDir() {
			newNode.Type = DIR
		} else {
			newNode.Type = FILE
		}

		dirNames = append(dirNames, &newNode)
	}
	return dirNames
}

func addJsonSuffix(name *string) {
	if !strings.HasSuffix(*name, ".json") {
		*name += ".json"
	}
}

// Not used anymore but was a fun exercice
func (ft *FileTreeExplorer) GetTheWholeTree() ([]*Node, error) {
	ft.FileTree = Node{
		Name:  "Lab",
		Type:  DIR,
		Files: make([]*Node, 0),
	}

	previousDepth := 0
	visited := make([]*Node, 0)
	lastInsertedNode := &ft.FileTree
	visited = append(visited, lastInsertedNode)

	err := filepath.WalkDir(ft.GetLabPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// On skip la lecture du dossier
		if path == ft.GetLabPath() {
			return nil
		}

		pathFromLab := strings.Split(path, ft.GetLabPath()+string(filepath.Separator))[1]
		nodes := strings.Split(pathFromLab, string(filepath.Separator))
		currentDepth := len(nodes)
		ft.Logger.Debug(pathFromLab + " Profondeur: " + fmt.Sprint(currentDepth))

		// Si on est allé à un noeud plus profond
		if previousDepth < currentDepth {
			n := visited[currentDepth-1]
			// Ajout du nouveau noeud dans le tableau des noeuds visités
			lastInsertedNode := n.InsertNode(d.IsDir(), nodes[currentDepth-1])
			visited = append(visited, lastInsertedNode)
			previousDepth = currentDepth
			return nil
		}

		// L'itération précédente était dans un noeud plus profond que l'actuelle.
		if previousDepth > currentDepth {
			n := visited[currentDepth-1]
			// On trie les noeuds de la profondeur précédente avant de la quitter
			n.SortNodes()
			// Création d'un nouveau noeud dans le noeud à la profondeur actuelle dans le tableau des noeuds visités
			lastInsertedNode = n.InsertNode(d.IsDir(), nodes[currentDepth-1])
			previousDepth = currentDepth
			// On modifie le tableau des noeuds visités pour supprimer les noeuds plus bas déjà visités
			visited = visited[:currentDepth]
			visited = append(visited, lastInsertedNode)
			return nil
		}

		// Si on est au même niveau que la boucle précédente
		if previousDepth == currentDepth {
			n := visited[currentDepth-1]
			// On trie les noeuds de la profondeur précédente avant de la quitter
			n.SortNodes()
			// Insertion du noeud à la profondeur
			lastInsertedNode = n.InsertNode(d.IsDir(), nodes[currentDepth-1])
			// On modifie le tableau des noeuds visités pour supprimer les noeuds plus bas déjà visités
			visited = visited[:currentDepth]
			visited = append(visited, lastInsertedNode)
			return nil
		}

		return nil
	})

	if err != nil {
		ft.Logger.Error("erreur lors de l'obtention des dossiers:" + err.Error() + " chemin: " + ft.GetLabPath())
		return nil, err
	}

	ft.FileTree.SortNodes()
	return ft.FileTree.Files, nil
}
