package filetree

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"flow-poc/backend/config"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type FileTreeExplorer struct {
	ctx    context.Context
	logger logger.Logger
	cfg    *config.AppConfig
}

type Node struct {
	files map[string]Node
}

func NewFileTree(cfg *config.AppConfig) *FileTreeExplorer {
	return &FileTreeExplorer{
		logger: logger.NewDefaultLogger(),
		cfg:    cfg,
	}
}

func (ft *FileTreeExplorer) SetContext(ctx context.Context) {
	ft.ctx = ctx
}

func (ft *FileTreeExplorer) SetConfigFile(cfg config.AppConfig) {
	ft.cfg = &cfg
}

func (ft *FileTreeExplorer) GetFileTree() (map[string]Node, error) {
	tree := make(map[string]Node)

	err := filepath.WalkDir(ft.cfg.ConfigFile.LabPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// On skip la lecture du dossier
		if path == ft.cfg.ConfigFile.LabPath {
			return nil
		}

		pathFromLab := strings.Split(path, ft.cfg.ConfigFile.LabPath+string(filepath.Separator))[1]
		nodes := strings.Split(pathFromLab, string(filepath.Separator))

		GoDownTree(d, tree, nodes)
		ft.logger.Debug(fmt.Sprintf("Chemin relatif au lab: %v", nodes))

		return nil
	})
	if err != nil {
		ft.logger.Error("erreur lors de l'obtention des dossiers:" + err.Error() + " chemin: " + ft.cfg.ConfigFile.LabPath)
		return nil, err
	}

	ft.logger.Debug(fmt.Sprintf("%v", tree))
	return tree, nil
}

func GoDownTree(d fs.DirEntry, tree map[string]Node, nodes []string) {
	// Noeud actuel puis noeuds restants
	currentNode, nextNode := nodes[0], nodes[1:]
	nodeName := ""

	if d.IsDir() {
		nodeName = currentNode
	} else {
		info, _ := d.Info()
		nodeName = info.Name()
	}

	_, ok := tree[nodeName]
	if !ok {
		tree[nodeName] = Node{
			files: make(map[string]Node),
		}
	}

	// Si les noeuds suivants n'existent pas, alors nous avont terminé la récursion
	if len(nextNode) == 0 {
		return
	}

	// Sinon on continue
	GoDownTree(d, tree[currentNode].files, nextNode)
}
