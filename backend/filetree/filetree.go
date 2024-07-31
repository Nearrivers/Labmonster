package filetree

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"flow-poc/backend/config"
)

var (
	ErrPathMissing = errors.New("the path must be specified")
)

type FileTreeExplorer struct {
	Cfg         *config.AppConfig
	Directories []string `json:"directories"`
}

func NewFileTree(cfg *config.AppConfig) *FileTreeExplorer {
	ft := &FileTreeExplorer{
		Cfg: cfg,
	}

	ft.GetLabDirs()

	return ft
}

func (ft *FileTreeExplorer) GetLabPath() string {
	return ft.Cfg.ConfigFile.LabPath
}

func (ft *FileTreeExplorer) GetDirectories() []string {
	ft.GetLabDirs()
	return ft.Directories
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

	nodes, err := createNodesFromDirEntries(entries)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// Takes an array of fs.DirEntry to create an array of type *Node and returns it
// This function ignore any .labmonster directory as it might contain config files that are not relevant to the user

// Get every directory name inside the lab and set the Directories
// of the FileTreeExplorer struct
func (ft *FileTreeExplorer) GetLabDirs() error {
	ft.Directories = make([]string, 0)

	ft.Directories = append(ft.Directories, "/")
	err := filepath.WalkDir(ft.GetLabPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() || d.Name() == ".labmonster" || path == ft.GetLabPath() {
			return nil
		}

		n := strings.TrimPrefix(path, ft.GetLabPath())
		ft.Directories = append(ft.Directories, strings.ReplaceAll(n[1:], string(filepath.Separator), "/"))
		return nil
	})

	return err
}

func createNodesFromDirEntries(entries []fs.DirEntry) ([]*Node, error) {
	dirNames := make([]*Node, 0)
	for _, entry := range entries {
		if entry.Name() == filepath.Ext(entry.Name()) {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		newNode := Node{
			Name:      strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())),
			Extension: filepath.Ext(entry.Name()),
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
