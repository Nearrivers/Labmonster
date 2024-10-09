// This package handles file operations inside a Lab.
package file_handler

import (
	"encoding/json"
	"errors"
	"flow-poc/backend/config"
	"flow-poc/backend/filesystem/node"
	"flow-poc/backend/filesystem/recentfiles"
	"flow-poc/backend/graph"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Maximum of recently opened files that the application remembers
	maxRecentlyOpenedFiles = 15
)

var (
	ErrFileAreDifferent   = errors.New("the 2 files are different")
	ErrEqualOldAndNewPath = errors.New("the old and paths must be different")
	ErrGetSubDirAndFile   = errors.New("can't build file tree")
	ErrNothingRead        = errors.New("nothing was read when trying to copy")
)

type FileHandler struct {
	// App's configuration
	Cfg         *config.AppConfig
	RecentFiles *recentfiles.RecentlyOpened
}

func NewFileHandler(cfg *config.AppConfig) *FileHandler {
	fh := &FileHandler{
		Cfg:         cfg,
		RecentFiles: recentfiles.NewRecentlyOpened(cfg, maxRecentlyOpenedFiles),
	}

	return fh
}

func (fh *FileHandler) GetLabPath() string {
	return fh.Cfg.ConfigFile.LabPath
}

func (fh *FileHandler) GetRecentlyOpenedFiles() ([]string, error) {
	return fh.RecentFiles.GetRecentlyOpenedFiles()
}

// Given a path to a directory starting from the lab root, this function will read
// its content using the os.ReadDir method, transforms those entries into Nodes
// and return them
func (fh *FileHandler) GetSubDirAndFiles(pathFromLabRoot string) ([]*node.Node, error) {
	dirPath := filepath.Join(fh.GetLabPath(), pathFromLabRoot)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, &GetSubDirAndFilesError{err}
	}

	nodes, err := node.CreateNodesFromDirEntries(entries)
	if err != nil {
		return nil, &GetSubDirAndFilesError{err}
	}

	return nodes, nil
}

func (fh *FileHandler) CreateFile(pathFromLabRoot string) (node.Node, error) {
	p := filepath.Join(fh.GetLabPath(), pathFromLabRoot)
	g := graph.GetInitGraph()

	if !doesFileExist(p) {
		f, err := os.Create(p)
		if err != nil {
			return node.Node{}, err
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			return node.Node{}, err
		}

		err = writeFile(g, f)
		if err != nil {
			return node.Node{}, err
		}

		n := node.NewNode(stat.Name(), ".json", node.FILE)
		return n, nil
	}

	f, name, err := createNonDuplicateFile(p)
	if err != nil {
		return node.Node{}, err
	}

	defer f.Close()

	err = writeFile(g, f)
	if err != nil {
		return node.Node{}, err
	}

	n := node.NewNode(name, ".json", node.FILE)
	return n, nil
}

func (fh *FileHandler) OpenFile(pathFromLabRoot string) (graph.Graph, error) {
	path := filepath.Join(fh.GetLabPath(), pathFromLabRoot)
	f, err := os.ReadFile(path)
	if err != nil {
		return graph.Graph{}, err
	}

	var g graph.Graph
	err = json.Unmarshal(f, &g)
	if err != nil {
		return graph.Graph{}, &OpenFileError{err}
	}

	fh.RecentFiles.AddRecentFile(pathFromLabRoot)

	return g, nil
}

func (fh *FileHandler) SaveFile(pathFromLabRoot string, graphToSave graph.Graph) error {
	path := filepath.Join(fh.GetLabPath(), pathFromLabRoot)
	if !doesFileExist(path) {
		return nil
	}

	// Create truncates the file if it already exists
	f, err := os.Create(path)
	if err != nil {
		return &SaveFileError{path, err}
	}
	defer f.Close()

	err = writeFile(graphToSave, f)
	if err != nil {
		return &SaveFileError{path, err}
	}

	return nil
}

// Rename a file on the user's machine
func (fh *FileHandler) RenameFile(pathFromRootOfTheLab, oldName, newName string) error {
	labPath := fh.GetLabPath()
	oldPath := filepath.Join(labPath, pathFromRootOfTheLab, oldName)
	newPath := filepath.Join(labPath, pathFromRootOfTheLab, newName)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	// TODO: Renommer l'entrée qui va avec dans les fichiers récents
	return nil
}

// Given the path to a file starting from the lab root,
// deletes a file on the user's machine and from the in-memory tree
func (fh *FileHandler) DeleteFile(pathFromRootOfTheLab string) error {
	err := os.Remove(filepath.Join(fh.GetLabPath(), pathFromRootOfTheLab))
	if err != nil {
		return err
	}

	fh.RecentFiles.RemoveRecent(pathFromRootOfTheLab)

	return nil
}

// Given a path to a file starting from the lab root and an another path to a directory,
// moves the file to the new directory.
func (fh *FileHandler) MoveFileToExistingDir(oldPath, newPath string) (string, error) {
	if oldPath == newPath {
		return "", ErrEqualOldAndNewPath
	}

	labPath := fh.GetLabPath()
	op := filepath.Join(labPath, oldPath)

	// If the file we want to move is at the root of the lab
	// that means oldPath = filename + file extension
	isPathLocatedAtLabRoot := !strings.Contains(oldPath, "/")
	if isPathLocatedAtLabRoot {
		if newPath == "/" {
			// If we're here that means the old path is the root of the lab.
			// That also means that if newPath == "/" and "/" being equivalent to root of the lab,
			// the old and new paths are equal
			return "", ErrEqualOldAndNewPath
		}

		np := filepath.Join(labPath, newPath, oldPath)
		return moveFile(op, np)
	}

	fileName := filepath.Base(oldPath)
	if oldPath == newPath+"/"+fileName {
		return "", ErrEqualOldAndNewPath
	}

	if newPath == "/" {
		np := filepath.Join(labPath, fileName)
		return fileName, os.Rename(op, np)
	}

	np := filepath.Join(labPath, newPath, fileName)
	return moveFile(op, np)
}

// Create a file named after the fileName argument. If the file already exists, it will try
// to add a number at the end to avoid duplicates
func (fh *FileHandler) DuplicateFile(pathToFileFromLabRoot, extension string) (newFileName string, error error) {
	labPath := fh.GetLabPath()
	path := filepath.Join(labPath, pathToFileFromLabRoot+extension)

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	f2, name, err := createNonDuplicateFile(path)
	if err != nil {
		return "", err
	}

	defer f.Close()
	defer f2.Close()

	n, err := io.Copy(f2, f)
	if err != nil {
		return "", err
	}

	if n == 0 {
		return "", ErrNothingRead
	}

	return name, nil
}

func writeFile(g graph.Graph, f *os.File) error {
	b, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		return &WriteFileError{f.Name(), err}
	}

	_, err = f.Write(b)
	if err != nil {
		return &WriteFileError{f.Name(), err}
	}

	return nil
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
