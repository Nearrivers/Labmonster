package dirhandler

import (
	"flow-poc/backend/config"
	"flow-poc/backend/filesystem/node"
	"flow-poc/backend/filesystem/recentfiles"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type DirHandler struct {
	Cfg         *config.AppConfig
	Directories []string `json:"directories"`
	recent      *recentfiles.RecentlyOpened
}

func NewDirHandler(cfg *config.AppConfig, recent *recentfiles.RecentlyOpened) *DirHandler {
	dh := &DirHandler{
		Cfg:    cfg,
		recent: recent,
	}

	return dh
}

func (dh *DirHandler) GetLabPath() string {
	return dh.Cfg.ConfigFile.LabPath
}

// Get every directory name inside the lab and set the Directories
// of the FileTreeExplorer struct
func (dh *DirHandler) GetLabDirs() error {
	dh.Directories = make([]string, 0)

	dh.Directories = append(dh.Directories, "/")
	err := filepath.WalkDir(dh.GetLabPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() || d.Name() == ".labmonster" || path == dh.GetLabPath() {
			return nil
		}

		n := strings.TrimPrefix(path, dh.GetLabPath())
		dh.Directories = append(dh.Directories, filepath.ToSlash(n))
		return nil
	})

	if err != nil {
		return &GetLabDirsError{err}
	}

	return nil
}

func (dh *DirHandler) GetDirectories() []string {
	dh.GetLabDirs()
	return dh.Directories
}

func (dh *DirHandler) CreateDirectory(pathFromLabRoot string) (node.Node, error) {
	p := filepath.Join(dh.GetLabPath(), pathFromLabRoot)

	if !doesDirExists(p) {
		err := os.Mkdir(p, os.ModeDir)
		if err != nil {
			return node.Node{}, err
		}

		name := filepath.Base(p)
		n := node.NewNode(name, "", node.DIR)
		return n, nil
	}

	name, dupErr := createNonDuplicateDir(p)
	if dupErr != nil {
		return node.Node{}, dupErr
	}

	n := node.NewNode(name, "", node.DIR)
	return n, nil
}

func (dh *DirHandler) DeleteDirectory(pathFromLabRoot string) error {
	p := filepath.Join(dh.GetLabPath(), pathFromLabRoot)

	err := os.RemoveAll(p)
	if err != nil {
		return err
	}

	dh.recent.CheckIfRecentFileStillExists()

	return nil
}

func (dh *DirHandler) RenameDirectory(oldPathFromRoot, newPathFromRoot string) error {
	labPath := dh.GetLabPath()
	p := filepath.Join(labPath, oldPathFromRoot)
	np := filepath.Join(labPath, newPathFromRoot)

	dh.recent.ReconcilePaths(oldPathFromRoot, newPathFromRoot)

	return os.Rename(p, np)
}

func doesDirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
