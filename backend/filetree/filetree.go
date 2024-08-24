package filetree

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"flow-poc/backend/config"
	"flow-poc/backend/filetree/recentfiles"
	"flow-poc/backend/graph"
)

const (
	maxRecentlyOpenedFiles = 15
)

var (
	ErrFileAreDifferent   = errors.New("the 2 files are different")
	ErrEqualOldAndNewPath = errors.New("the old and paths must be different")
	ErrGetSubDirAndFile   = errors.New("can't build file tree")
	ErrNothingRead        = errors.New("nothing was read when trying to copy")
)

type FileTree struct {
	Cfg         *config.AppConfig
	Directories []string `json:"directories"`
	RecentFiles recentfiles.RecentlyOpened
}

func NewFileTree(cfg *config.AppConfig) *FileTree {
	ft := &FileTree{
		Cfg:         cfg,
		RecentFiles: *recentfiles.NewRecentlyOpened(cfg, maxRecentlyOpenedFiles),
	}

	ft.GetLabDirs()

	return ft
}

func (ft *FileTree) GetLabPath() string {
	return ft.Cfg.ConfigFile.LabPath
}

func (ft *FileTree) GetDirectories() []string {
	ft.GetLabDirs()
	return ft.Directories
}

func (ft *FileTree) GetRecentlyOpenedFiles() ([]string, error) {
	return ft.RecentFiles.GetRecentlyOpenedFiles()
}

// Given a path to a directory starting from the lab root, this function will read
// its content using the os.ReadDir method, transforms those entries into Nodes
// and return them
func (ft *FileTree) GetSubDirAndFiles(pathFromLabRoot string) ([]*Node, error) {
	dirPath := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, &GetSubDirAndFilesError{err}
	}

	nodes, err := createNodesFromDirEntries(entries)
	if err != nil {
		return nil, &GetSubDirAndFilesError{err}
	}

	return nodes, nil
}

// Takes an array of fs.DirEntry to create an array of type *Node and returns it
// This function ignore any .labmonster directory as it might contain config files that are not relevant to the user

// Get every directory name inside the lab and set the Directories
// of the FileTreeExplorer struct
func (ft *FileTree) GetLabDirs() error {
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
		ft.Directories = append(ft.Directories, filepath.ToSlash(n))
		return nil
	})

	if err != nil {
		return &GetLabDirsError{err}
	}

	return nil
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (ft *FileTree) CreateFile(pathFromLabRoot string) (Node, error) {
	p := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	g := graph.GetInitGraph()

	if !doesFileExist(p) {
		f, err := os.Create(p)
		if err != nil {
			return Node{}, err
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			return Node{}, err
		}

		err = writeFile(g, f)
		if err != nil {
			return Node{}, err
		}

		n := NewNode(stat.Name(), ".json", FILE)
		return n, nil
	}

	f, name, err := createNonDuplicateFile(p)
	if err != nil {
		return Node{}, err
	}

	defer f.Close()

	err = writeFile(g, f)
	if err != nil {
		return Node{}, err
	}

	n := NewNode(name, ".json", FILE)
	return n, nil
}

func (ft *FileTree) OpenFile(pathFromLabRoot string) (graph.Graph, error) {
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	f, err := os.ReadFile(path)
	if err != nil {
		return graph.Graph{}, err
	}

	var g graph.Graph
	err = json.Unmarshal(f, &g)
	if err != nil {
		return graph.Graph{}, &OpenFileError{err}
	}

	ft.RecentFiles.AddRecentFile(pathFromLabRoot)

	return g, nil
}

// Sauvegarde le fichier JSON du graph
func (ft *FileTree) SaveFile(pathFromLabRoot string, graphToSave graph.Graph) error {
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
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
func (ft *FileTree) RenameFile(pathFromRootOfTheLab, oldName, newName string) error {
	labPath := ft.GetLabPath()
	oldPath := filepath.Join(labPath, pathFromRootOfTheLab, oldName)
	newPath := filepath.Join(labPath, pathFromRootOfTheLab, newName)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	return nil
}

// Given the path to a file starting from the lab root,
// deletes a file on the user's machine and from the in-memory tree
func (ft *FileTree) DeleteFile(pathFromRootOfTheLab string) error {
	err := os.Remove(filepath.Join(ft.GetLabPath(), pathFromRootOfTheLab))
	if err != nil {
		return err
	}

	ft.RecentFiles.RemoveRecent(pathFromRootOfTheLab)

	return nil
}

// Given a path to a file starting from the lab root and an another path to a directory,
// moves the file to the new directory.
func (ft *FileTree) MoveFileToExistingDir(oldPath, newPath string) (string, error) {
	if oldPath == newPath {
		return "", ErrEqualOldAndNewPath
	}

	labPath := ft.GetLabPath()
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
func (ft *FileTree) DuplicateFile(pathToFileFromLabRoot, extension string) (newFileName string, error error) {
	labPath := ft.GetLabPath()
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

func moveFile(oldPath, newPath string) (string, error) {
	if !doesFileExist(newPath) {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return "", nil
		}

		info, err := os.Stat(newPath)
		if err != nil {
			return "", err
		}
		return info.Name(), nil
	}

	oldFile, err := os.Open(oldPath)
	if err != nil {
		return "", err
	}
	defer oldFile.Close()

	newFile, name, err := createNonDuplicateFile(newPath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	_, err = io.Copy(oldFile, newFile)
	return name, err
}

// Given an absolute path to a file we're trying to create. Before calling this function, we know that this file
// already exists. This function will try to create the same file but with a number appended to its name in order to avoid duplicates.
// It will continue to increment the appended number as long as the function finds a file with the same name. After succeeding in creating the file,
// it will return an os.File pointer to it that the caller will have to close, the actual name of the file to avoid f.Stat() boilerplate and an error
func createNonDuplicateFile(absPath string) (*os.File, string, error) {
	p := filepath.Dir(absPath)
	b := filepath.Base(absPath)
	newFileName := b[:strings.LastIndex(b, ".")]
	ext := filepath.Ext(absPath)

	for i := 1; ; i++ {
		name := fmt.Sprintf("%s %d%s", newFileName, i, ext)
		if doesFileExist(filepath.Join(p, name)) {
			continue
		}

		f, err := os.Create(filepath.Join(p, name))
		if err != nil {
			return nil, "", err
		}

		return f, name, nil
	}
}

func createNodesFromDirEntries(entries []fs.DirEntry) ([]*Node, error) {
	dirNames := make([]*Node, 0)
	for _, entry := range entries {
		ext := filepath.Ext(entry.Name())
		if entry.Name() == ext {
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
			newNode.FileType = DetectFileType(ext)
			newNode.Type = FILE
		}

		dirNames = append(dirNames, &newNode)
	}
	return dirNames, nil
}

// Given an extension, it wil return the corresponding FileType
func DetectFileType(extension string) FileType {
	switch extension {
	case ".png", ".jpeg", ".gif", ".webp":
		return IMAGE
	case ".json":
		return GRAPH
	case ".mp4", ".mpeg":
		return VIDEO
	default:
		return UNSUPPORTED
	}
}
