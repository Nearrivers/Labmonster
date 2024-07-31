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
	ErrFileAreDifferent   = errors.New("the 2 files are different")
	ErrEqualOldAndNewPath = errors.New("the old and paths must be different")
)

// Short function that, given an absolute path to a file, tells wheter or not the file exists
func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (ft *FileTreeExplorer) CreateFile(pathFromLabRoot string) (Node, error) {
	p := filepath.Join(ft.GetLabPath(), pathFromLabRoot)

	if !doesFileExist(p) {
		f, err := os.Create(p)
		if err != nil {
			return Node{}, nil
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			return Node{}, nil
		}

		n := NewNode(stat.Name(), "json", FILE)
		return n, nil
	}

	f, name, err := createNonDuplicateFile(p)
	if err != nil {
		return Node{}, nil
	}

	defer f.Close()
	n := NewNode(name, "json", FILE)
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

		n := NewNode(newFileName, "json", FILE)
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

		n := NewNode(name, "json", FILE)
		return n, nil
	}
}

// Sauvegarde le fichier JSON du graph
func (ft *FileTreeExplorer) SaveFile() {}

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
	if !strings.Contains(oldPath, "/") {
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
	newFileName := absPath[strings.LastIndex(absPath, string(filepath.Separator))+1:strings.LastIndex(absPath, ".")]
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

// Function used in tests
// Compare 2 files content. The goal is to fail asap
// Read files by chunks defined by the chunkSize variable
func (ft *FileTreeExplorer) fileContentCompare(file1, file2 string) error {
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

	statf2, err := f2.Stat()
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
