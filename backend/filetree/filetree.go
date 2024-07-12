package filetree

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"flow-poc/backend/config"

	"github.com/wailsapp/wails/v2/pkg/logger"
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

func (ft *FileTreeExplorer) GetFirstDepth() ([]*Node, error) {
	entries, err := os.ReadDir(ft.GetLabPath())
	if err != nil {
		return nil, err
	}

	dirNames := make([]*Node, 0)

	for _, entry := range entries {
		if entry.Name() == ".labmonster" && entry.IsDir() {
			continue
		}

		newSimpleNode := Node{
			Name:  entry.Name(),
			Files: make([]*Node, 0),
		}

		if entry.IsDir() {
			newSimpleNode.Type = DIR
		} else {
			newSimpleNode.Type = FILE
		}
		dirNames = append(dirNames, &newSimpleNode)
	}

	SortNodes(dirNames)
	return dirNames, nil
}

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

	// TODO: Cette méthode peut très probablement être optimisée
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
