package filetree

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrNoFileInThisLevel = errors.New("no file at this level")
	ErrNodeNotFound      = errors.New("file not found")
	ErrFileAreDifferent  = errors.New("the 2 files are different")
)

// Fonction utilitaire qui permet de déterminer si un fichier existe au chemin indiqué
func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Créer un fichier à la racine et l'ajoute aux noeuds
func (ft *FileTreeExplorer) CreateNewFileAtRoot(newFileName string) (string, error) {
	// Création d'un fichier si ce dernier n'existe pas
	if !doesFileExist(filepath.Join(ft.GetLabPath(), newFileName+".json")) {
		f, err := os.Create(filepath.Join(ft.GetLabPath(), newFileName+".json"))
		if err != nil {
			return "", err
		}

		defer f.Close()

		newNode := InsertNode(false, &ft.FileTree, newFileName+".json")
		SortNodes(ft.FileTree.Files)
		return newNode.Name, nil
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
			return "", err
		}

		defer f.Close()

		InsertNode(false, &ft.FileTree, name)
		SortNodes(ft.FileTree.Files)
		return name, nil
	}
}

// Sauvegarde le fichier JSON du graph
func (ft *FileTreeExplorer) SaveFile() {}

// Rename a file on the user's machine and inside the in-memory tree
func (ft *FileTreeExplorer) RenameFile(pathFromRootOfTheLab, oldName, newName string) error {
	labPath := ft.GetLabPath()
	oldPath := filepath.Join(labPath, pathFromRootOfTheLab, oldName)
	newPath := filepath.Join(labPath, pathFromRootOfTheLab, newName+".json")

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	// TODO: Renommer le fichier dans l'arbre en mémoire

	return nil
}

// Deletes a file on the user's machine and from the in-memory tree
func (ft *FileTreeExplorer) DeleteFile(pathFromRootOfTheLab string) error {
	if !strings.Contains(string(pathFromRootOfTheLab), ".json") {
		pathFromRootOfTheLab += ".json"
	}

	err := os.Remove(filepath.Join(ft.GetLabPath(), pathFromRootOfTheLab))
	if err != nil {
		return err
	}

	// TODO: Suppression du fichier dans l'arbre en mémoire

	return nil
}

// Déplace le fichier
func (ft *FileTreeExplorer) MoveFile(oldPath, newPath string) {

}

// Create a file named after the fileName argument. If the file already exists, it will try
// to add a number at the end to avoid duplicates
func (ft *FileTreeExplorer) DuplicateFile(pathToFileFromLabRoot, fileName string) (newFileName string, error error) {
	labPath := ft.GetLabPath()
	path := filepath.Join(labPath, pathToFileFromLabRoot)
	old := filepath.Join(path, fileName+".json")

	f, err := os.Open(old)
	if err != nil {
		return "", fmt.Errorf("can't open file: %v", err)
	}

	defer f.Close()

	var f2 *os.File
	var newFile string

	for i := 1; ; i++ {
		newFile = filepath.Join(path, fmt.Sprintf("%s %d.json", fileName, i))
		if !doesFileExist(newFile) {
			f2, err = os.Create(newFile)
			if err != nil {
				return "", err
			}

			defer f2.Close()

			break
		}
	}

	err = copyFile(f, f2)
	if err != nil {
		return "", err
	}

	stat, err := f2.Stat()
	if err != nil {
		return "", err
	}

	return stat.Name(), nil
}

// Take two files and copy the content of the first into the second
func copyFile(inputFile, outputFile *os.File) error {
	_, err := io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

// Function used in tests
// Compare 2 files content. The goal is to fail asap
// Read files by chunks defined by the chunkSize variable
func fileContentCompare(ft *FileTreeExplorer, file1, file2 string) error {
	const chunkSize = 2000
	path1 := filepath.Join(ft.GetLabPath(), file1)
	path2 := filepath.Join(ft.GetLabPath(), file2)

	f1, err := os.Open(path1)
	if err != nil {
		return fmt.Errorf("couldn't open the first file: %v", err)
	}
	defer f1.Close()

	f2, err := os.Open(path2)
	if err != nil {
		return fmt.Errorf("couldn't open the second file: %v", err)
	}
	defer f2.Close()

	statf1, err := f1.Stat()
	if err != nil {
		return fmt.Errorf("couldn't get the first file info: %v", err)
	}

	statf2, err := f1.Stat()
	if err != nil {
		return fmt.Errorf("couldn't get the second file info: %v", err)
	}

	if statf1.Size() != statf2.Size() {
		return ErrFileAreDifferent
	}

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil && err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return nil
			}

			if err1 == io.EOF || err2 == io.EOF {
				return ErrFileAreDifferent
			}
		}

		if !bytes.Equal(b1, b2) {
			return ErrFileAreDifferent
		}
	}
}

// Binary search of a node using its name
// Node names are unique and they are sorted in alphabetical order
func searchFileOrDir(name string, level []*Node) (*Node, int, error) {
	if len(level) == 0 {
		return nil, -1, ErrNoFileInThisLevel
	}

	if len(level) == 1 {
		if level[0].Name == name {
			return level[0], 0, nil
		} else {
			return nil, -1, ErrNodeNotFound
		}
	}

	low := 0
	high := len(level) - 1

	for low <= high {
		median := (low + high) / 2
		medianName := level[median]

		if medianName.Name < name {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low != len(level) && level[low].Name == name {
		return level[low], low, nil
	}

	return nil, -1, ErrNodeNotFound
}
