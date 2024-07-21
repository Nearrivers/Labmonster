package filetree

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"flow-poc/backend/config"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

var (
	ErrPathMissing = errors.New("the path must be specified")
)

type FileTreeExplorer struct {
	Logger logger.Logger     `json:"logger"`
	Cfg    *config.AppConfig `json:"cfg"`
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

	nodes, err := ft.createNodesFromDirEntries(entries)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// Takes an array of fs.DirEntry to create an array of type *Node and returns it
// This function ignore any .labmonster directory as it might contain config files that are not relevant to the user
func (ft *FileTreeExplorer) createNodesFromDirEntries(entries []fs.DirEntry) ([]*Node, error) {
	dirNames := make([]*Node, 0)
	for _, entry := range entries {
		if entry.Name() == ".labmonster" && entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		newNode := Node{
			Name:      entry.Name(),
			Files:     make([]*Node, 0),
			CreatedAt: time.Now(),
			UpdatedAt: info.ModTime(),
		}

		if entry.IsDir() {
			newNode.Type = DIR
		} else {
			newNode.Type = FILE
		}

		dirNames = append(dirNames, &newNode)
	}
	return dirNames, nil
}
