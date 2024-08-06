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
	"flow-poc/backend/graph"
)

var (
	ErrFileAreDifferent   = errors.New("the 2 files are different")
	ErrEqualOldAndNewPath = errors.New("the old and paths must be different")
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

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (ft *FileTreeExplorer) CreateFile(pathFromLabRoot string) (Node, error) {
	p := filepath.Join(ft.GetLabPath(), pathFromLabRoot)

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

		n := NewNode(stat.Name(), ".json", FILE)
		return n, nil
	}

	f, name, err := createNonDuplicateFile(p)
	if err != nil {
		return Node{}, err
	}

	defer f.Close()
	n := NewNode(name, ".json", FILE)
	return n, nil
}

// Create a file at lab root
func (ft *FileTreeExplorer) CreateNewFileAtRoot(newFileName string) (Node, error) {
	// Création d'un fichier si ce dernier n'existe pas
	if !doesFileExist(filepath.Join(ft.GetLabPath(), newFileName+".json")) {
		f, err := os.Create(filepath.Join(ft.GetLabPath(), newFileName+".json"))
		if err != nil {
			return Node{}, err
		}

		defer f.Close()

		n := NewNode(newFileName, ".json", FILE)
		return n, nil
	}

	// Si un fichier avec le même nom existe déjà alors on va boucler avec un compteur afin de créer un
	// "Sans titre n.json" avec n un compteur incrémenté à chaque fois qu'un fichier du même nom est trouvé
	cpt := 1
	for {
		name := fmt.Sprintf("%s %d.json", newFileName, cpt)
		if doesFileExist(filepath.Join(ft.GetLabPath(), name)) {
			cpt++
			continue
		}

		f, err := os.Create(filepath.Join(ft.GetLabPath(), name))
		if err != nil {
			return Node{}, err
		}

		defer f.Close()

		n := NewNode(name, ".json", FILE)
		return n, nil
	}
}

func (ft *FileTreeExplorer) OpenFile(pathFromLabRoot string) (graph.Graph, error) {
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	f, err := os.ReadFile(path)
	if err != nil {
		return graph.Graph{}, err
	}

	var g graph.Graph
	err = json.Unmarshal(f, &g)
	if err != nil {
		return graph.Graph{}, err
	}

	return g, nil
}

// Sauvegarde le fichier JSON du graph
func (ft *FileTreeExplorer) SaveFile(pathFromLabRoot string, graphToSave graph.Graph) error {
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(graphToSave)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}

// Rename a file on the user's machine and inside the in-memory tree
func (ft *FileTreeExplorer) RenameFile(pathFromRootOfTheLab, oldName, newName string) error {
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
// The function will add the ".json" file extension if it's missing
// from the path
func (ft *FileTreeExplorer) DeleteFile(pathFromRootOfTheLab string) error {
	err := os.Remove(filepath.Join(ft.GetLabPath(), pathFromRootOfTheLab))
	if err != nil {
		return err
	}

	return nil
}

// Given a path to a file starting from the lab root and an another path to a directory,
// moves the file to the new directory.
func (ft *FileTreeExplorer) MoveFileToExistingDir(oldPath, newPath string) (string, error) {
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

	fileName := oldPath[strings.LastIndex(oldPath, "/")+1:]
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
func (ft *FileTreeExplorer) DuplicateFile(pathToFileFromLabRoot, extension string) (newFileName string, error error) {
	labPath := ft.GetLabPath()
	path := filepath.Join(labPath, pathToFileFromLabRoot+extension)

	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("can't open file: %v", err)
	}

	f2, name, err := createNonDuplicateFile(path)
	if err != nil {
		return "", err
	}

	defer f.Close()
	defer f2.Close()

	_, err = io.Copy(f, f2)
	if err != nil {
		return "", err
	}

	return name, nil
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
	p := absPath[:strings.LastIndex(absPath, string(filepath.Separator))+1]
	newFileName := absPath[strings.LastIndex(absPath, string(filepath.Separator))+1 : strings.LastIndex(absPath, ".")]
	ext := filepath.Ext(absPath)

	for i := 1; ; i++ {
		name := fmt.Sprintf("%s %d%s", newFileName, i, ext)
		if doesFileExist(filepath.Join(p, name)) {
			i++
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
