package filetree

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"flow-poc/backend/config"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type NodeType string

const (
	FILE NodeType = "FILE"
	DIR  NodeType = "DIR"
)

type FileTreeExplorer struct {
	Ctx      context.Context   `json:"ctx"`
	Logger   logger.Logger     `json:"logger"`
	Cfg      *config.AppConfig `json:"cfg"`
	FileTree Node              `json:"file_tree"`
}

type Node struct {
	Name  string   `json:"name"`
	Type  NodeType `json:"type"`
	Files []*Node  `json:"files"`
}

func NewFileTree(cfg *config.AppConfig) *FileTreeExplorer {
	return &FileTreeExplorer{
		Logger: logger.NewDefaultLogger(),
		Cfg:    cfg,
	}
}

func (ft *FileTreeExplorer) SetContext(ctx context.Context) {
	ft.Ctx = ctx
}

func (ft *FileTreeExplorer) SetConfigFile(cfg config.AppConfig) {
	ft.Cfg = &cfg
}

func (ft *FileTreeExplorer) GetFileTree() ([]*Node, error) {
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
	err := filepath.WalkDir(ft.Cfg.ConfigFile.LabPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// On skip la lecture du dossier
		if path == ft.Cfg.ConfigFile.LabPath {
			return nil
		}

		pathFromLab := strings.Split(path, ft.Cfg.ConfigFile.LabPath+string(filepath.Separator))[1]
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
		ft.Logger.Error("erreur lors de l'obtention des dossiers:" + err.Error() + " chemin: " + ft.Cfg.ConfigFile.LabPath)
		return nil, err
	}

	SortNodes(ft.FileTree.Files)
	return ft.FileTree.Files, nil
}

func SortNodes(files []*Node) {
	sort.SliceStable(files, func(i, j int) bool {
		if files[i].Type != files[j].Type {
			return files[i].Type < files[j].Type
		}

		return files[i].Name < files[j].Name
	})
}

func InsertNode(isDir bool, node *Node, name string) *Node {
	var nodetype NodeType

	if isDir {
		nodetype = DIR
	} else {
		nodetype = FILE
	}

	newNode := Node{
		Name:  name,
		Type:  nodetype,
		Files: []*Node{},
	}

	node.Files = append(node.Files, &newNode)

	return &newNode
}
