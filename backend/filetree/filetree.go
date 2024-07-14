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

func (ft *FileTreeExplorer) GetSubDirAndFiles(pathFromLabRoot string) ([]*Node, error) {
	dirPath := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	nodes := ft.createNodesFromDirEntries(entries)

	if len(nodes) > 0 {
		err = ft.SetNodeFiles(pathFromLabRoot, nodes)
		if err != nil {
			return nil, err
		}
	}

	return nodes, nil
}

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

func (ft *FileTreeExplorer) RemoveNode(pathFromLabRoot string) error {
	if pathFromLabRoot == "" {
		return ErrPathMissing
	}

	path := strings.Split(pathFromLabRoot, string(filepath.Separator))

	if len(path) == 1 {
		_, i, err := searchFileOrDir(pathFromLabRoot, ft.FileTree.Files)
		if err != nil {
			return err
		}

		newFiles, err := removeIndex(ft.FileTree.Files, i)
		if err != nil {
			return err
		}

		ft.FileTree.Files = newFiles
		return nil
	}

	n, i, err := ft.FindNodeWithPath(pathFromLabRoot)
	if err != nil {
		return err
	}

	newFiles, err := removeIndex(n.Files, i)
	if err != nil {
		return err
	}

	n.Files = newFiles
	return nil
}

func (ft *FileTreeExplorer) FindNodeWithPath(pathFromLabRoot string) (*Node, int, error) {
	separator := string(filepath.Separator)
	path := strings.Split(pathFromLabRoot, separator)
	n := &ft.FileTree
	index := 0

	for _, dir := range path {
		SortNodes(n.Files)
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
			// Ajout du nouveau noeud dans le tableau des noeuds visités
			lastInsertedNode := InsertNode(d.IsDir(), visited[currentDepth-1], nodes[currentDepth-1])
			visited = append(visited, lastInsertedNode)
			previousDepth = currentDepth
			return nil
		}

		// L'itération précédente était dans un noeud plus profond que l'actuelle.
		if previousDepth > currentDepth {
			// On trie les noeuds de la profondeur précédente avant de la quitter
			SortNodes(visited[currentDepth-1].Files)
			// Création d'un nouveau noeud dans le noeud à la profondeur actuelle dans le tableau des noeuds visités
			lastInsertedNode = InsertNode(d.IsDir(), visited[currentDepth-1], nodes[currentDepth-1])
			previousDepth = currentDepth
			// On modifie le tableau des noeuds visités pour supprimer les noeuds plus bas déjà visités
			visited = visited[:currentDepth]
			visited = append(visited, lastInsertedNode)
			return nil
		}

		// Si on est au même niveau que la boucle précédente
		if previousDepth == currentDepth {
			// On trie les noeuds de la profondeur précédente avant de la quitter
			SortNodes(visited[previousDepth-1].Files)
			// Insertion du noeud à la profondeur
			lastInsertedNode = InsertNode(d.IsDir(), visited[currentDepth-1], nodes[currentDepth-1])
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

	SortNodes(ft.FileTree.Files)
	return ft.FileTree.Files, nil
}
