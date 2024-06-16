package filetree

import (
	"context"
	"io/fs"
	"path/filepath"
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

		CreateFileTree(d, &ft.FileTree, nodes)
		return nil
	})

	ft.PrintTree()
	if err != nil {
		ft.Logger.Error("erreur lors de l'obtention des dossiers:" + err.Error() + " chemin: " + ft.Cfg.ConfigFile.LabPath)
		return nil, err
	}

	return ft.FileTree.Files, nil
}

func CreateFileTree(d fs.DirEntry, node *Node, nodeNames []string) {
	// Noeud actuel puis noeuds restants
	currentNodeName, nextNodeNames := nodeNames[0], nodeNames[1:]

	// Si l'arbre est vide, on insère le premier noeud d'office
	if len(node.Files) == 0 {
		var nodetype NodeType

		if d.IsDir() {
			nodetype = DIR
		} else {
			nodetype = FILE
		}

		node.Files = append(node.Files, &Node{
			Name:  currentNodeName,
			Type:  nodetype,
			Files: []*Node{},
		})
	}

	// Si les noeuds suivants n'existent pas, alors nous avont terminé la récursion
	if len(nextNodeNames) == 0 {
		return
	}

	var nextNode *Node
	for _, n := range node.Files {
		if n.Name == currentNodeName {
			nextNode = n
		}
	}

	if nextNode == nil {
		newNode := InsertNode(d.IsDir(), node, currentNodeName)
		nextNode = &newNode
	}

	CreateFileTree(d, nextNode, nextNodeNames)
}

func InsertNode(isDir bool, node *Node, name string) Node {
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

	return newNode
}
