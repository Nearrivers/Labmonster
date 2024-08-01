package filetree

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"flow-poc/backend/config"
)

// Creates dir by joining the last 2 args with filepath.Join. The first
// string arg should be the path to the temp dir and the second the path
// to the new dir. os.MkDirAll is used so it will create any dir necessary
// to reach the path given in the last arg
func createDirHelper(t testing.TB, tempDirPath, dirName string) {
	t.Helper()

	err := os.MkdirAll(filepath.Join(tempDirPath, dirName), 0750)
	if err != nil {
		t.Fatalf("couldn't create the first subdirectory: %v", err)
	}
}

// Creates temporary dir, sets it as lab's root and creates a file at root.
// Returns the path to the temporary dir and the filetree that resulted.
// The function os.RemoveAll(dir) should be called with the defer keyword
// to clean up after tests that use this helper function
func createTempDir(t testing.TB, testDirName, testFileName string) (string, *FileTreeExplorer) {
	t.Helper()

	tempDir := os.TempDir()

	dir, err := os.MkdirTemp("", testDirName)
	if err != nil {
		t.Fatalf("an error occured while creating temporary directory: %v %s", err, tempDir)
	}

	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	})

	ft.CreateNewFileAtRoot(testFileName)

	return dir, ft
}

func TestGetSubDirAndFiles(t *testing.T) {
	t.Run("read first level", func(t *testing.T) {
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testLab", subFile1)
		defer os.RemoveAll(dir)
		subDir1 := "testDir1"
		createDirHelper(t, dir, subDir1)

		nodes, err := ft.GetSubDirAndFiles("")
		if err != nil {
			t.Fatalf("couldn't get first tree depth: %v", err)
		}

		if len(nodes) != 2 {
			t.Errorf("want 2 nodes, got %d", len(nodes))
		}

		if nodes[0].Name != subDir1 {
			t.Errorf("the first node should be the directory, got %s", nodes[0].Name)
		}
	})

	t.Run("read next level", func(t *testing.T) {
		subDir1 := "testDir1"
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testNextLevel", subFile1)
		defer os.RemoveAll(dir)
		createDirHelper(t, dir, subDir1)

		subDir2 := "testSubDir"
		createDirHelper(t, filepath.Join(dir, subDir1), subDir2)

		nodes, err := ft.GetSubDirAndFiles(subDir1)
		if err != nil {
			t.Fatalf("couldn't get first tree depth: %v", err)
		}

		if len(nodes) != 1 {
			t.Errorf("want 1 node, got %d", len(nodes))
		}

		if nodes[0].Name != subDir2 {
			t.Errorf("the first node should be the directory, got %s", nodes[0].Name)
		}
	})

	// TODO: Ajouter d'autres tests dans d'autres profondeurs lorsque la fonction de création sera
	// implémentée
}

func TestGetLabDirs(t *testing.T) {
	t.Run("get every directory of the lab", func(t *testing.T) {
		subDir1 := "testDir1"
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testNextLevel", subFile1)
		defer os.RemoveAll(dir)
		createDirHelper(t, dir, subDir1)

		subDir2 := "testSubDir"
		createDirHelper(t, filepath.Join(dir, subDir1), subDir2)

		err := ft.GetLabDirs()
		if err != nil {
			t.Fatalf("couldn't get first tree depth: %v", err)
		}

		if len(ft.Directories) != 3 {
			t.Errorf("want 3 directories, got %d with %s", len(ft.Directories), strings.Join(ft.Directories, ", "))
		}
	})
}


func getNewFileTreeExplorer() (*FileTreeExplorer, string) {
	dir, _ := os.MkdirTemp("", "testFt")
	return NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	}), dir
}

func createFileBeforeTest(t testing.TB, ft *FileTreeExplorer, fileName string) {
	t.Helper()

	_, err := ft.CreateNewFileAtRoot(fileName)
	if err != nil {
		t.Errorf("Error when creating the file: %v", err)
	}
}

func TestCreateNewFile(t *testing.T) {
	t.Run("Creating one file", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		testFileName := "happyPath test.json"
		defer ft.DeleteFile(testFileName)

		got, err := ft.CreateFile(testFileName)
		if err != nil {
			t.Fatalf("An error occured while creating the file: %v", err.Error())
		}

		assertFileExistence(t, ft.GetLabPath(), got.Name)
	})

	t.Run("Creating multiple files in a row", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		testFileName := "Multiple test.json"

		cpt := 0
		for cpt < 3 {
			newFile, err := ft.CreateFile(testFileName)
			if err != nil {
				t.Errorf("An error occured while creating the file: %s", err.Error())
			}
			defer ft.DeleteFile(newFile.Name)

			assertFileExistence(t, ft.GetLabPath(), newFile.Name)
			cpt++
		}
	})

	t.Run("Creating a file not at root", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		testDirPath := filepath.Join(dir, "create", "test")
		err := os.MkdirAll(testDirPath, os.ModePerm)
		if err != nil {
			t.Fatalf("An error occured while creating dirs for the test: %v", err)
		}

		testFileName := "create/test/deep create test.json"
		newFile, err := ft.CreateFile(testFileName)
		if err != nil {
			t.Errorf("An error occured while creating the file: %s", err.Error())
		}
		defer ft.DeleteFile(newFile.Name)

		assertFileExistence(t, ft.GetLabPath(),"create", "test", newFile.Name)
	})
}

func TestRenameFile(t *testing.T) {
	t.Run("Existing file rename at first level", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		oldName := "test rename"
		newName := "test rename 1.json"

		createFileBeforeTest(t, ft, oldName)
		defer ft.DeleteFile(newName)

		err := ft.RenameFile("", oldName+".json", newName)
		if err != nil {
			t.Fatalf("got an error %v", err)
		}
	
		assertFileExistence(t, ft.GetLabPath(), newName)
	})

	t.Run("Non existant file rename at first level", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		fakeName := "test fake rename"
		oldName := "test bad rename"
		newName := "test fake rename 2"

		createFileBeforeTest(t, ft, fakeName)

		defer ft.DeleteFile(fakeName)

		got := ft.RenameFile("", oldName+".json", newName)
		want := os.ErrNotExist

		if !errors.Is(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestDeleteFile(t *testing.T) {
	t.Run("Existing file deletion at first level", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		fileName := "test delete"

		createFileBeforeTest(t, ft, fileName)

		err := ft.DeleteFile(fileName+".json")
		if err != nil {
			t.Errorf("An error occured while deleting the file: %v", err)
		}

		want := false
		got := doesFileExist(filepath.Join(ft.GetLabPath(), fileName+".json"))

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Non existant file deletion at first level", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		fileName := "test non existant"

		createFileBeforeTest(t, ft, "non-existant")

		defer ft.DeleteFile("non-existant")

		got := ft.DeleteFile(fileName)
		want := os.ErrNotExist

		if got == nil {
			t.Error("An occured didn't occur but should have")
		}

		if !errors.Is(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertSameFiles(t testing.TB, ft *FileTreeExplorer, file1, file2 string) {
	t.Helper()

	if err := fileContentCompare(ft.GetLabPath(), file1, file2); err != nil {
		if err == ErrFileAreDifferent {
			t.Fatalf("the two files are different: %v", err)
		}

		t.Errorf("an error occured before comparing the two files: %v", err)
	}
}

func TestDuplicateFile(t *testing.T) {
	t.Run("File duplication at first level", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		fileName := "duplication test"
		createFileBeforeTest(t, ft, fileName)
		defer ft.DeleteFile(fileName)

		duplicatedFile, err := ft.DuplicateFile(fileName, ".json")
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}
		defer ft.DeleteFile(duplicatedFile)

		assertSameFiles(t, ft, fileName+".json", duplicatedFile)
	})

	t.Run("Running multiple file duplications in a row", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		fileName := "multiple duplication test"
		createFileBeforeTest(t, ft, fileName)
		defer ft.DeleteFile(fileName)

		duplicatedFile1, err := ft.DuplicateFile(fileName, ".json")
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}
		defer ft.DeleteFile(duplicatedFile1)

		duplicatedFile2, err := ft.DuplicateFile(fileName, ".json")
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}
		defer ft.DeleteFile(duplicatedFile2)

		assertSameFiles(t, ft, fileName+".json", duplicatedFile1)
		assertSameFiles(t, ft, fileName+".json", duplicatedFile2)
	})

	t.Run("Trying to duplicate a file that doesn't exists", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)
		fileName := "fake duplication test"

		_, err := ft.DuplicateFile(fileName, ".json")
		if err == nil {
			t.Fatal("didn't get an error, but wanted one")
		}
	})
}

func createFileHelper(t testing.TB, tempDirPath, completeFileName string) {
	t.Helper()

	f, err := os.Create(filepath.Join(tempDirPath, completeFileName))
	if err != nil {
		t.Fatalf("could not create file: %v", err)
	}
	f.Close()
}

func assertFileExistence(t testing.TB, path ...string) {
	t.Helper()

	if !doesFileExist(strings.Join(path, string(filepath.Separator))) {
		t.Error("the file was not moved")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("wrong error, got %v, want %v", got, want)
	}
}

func TestMoveFile(t *testing.T) {
	t.Run("move a file from lab root to another directory", func(t *testing.T) {
		subDir1 := "testDir1"
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testMoveFromRoot", subFile1)
		defer os.RemoveAll(dir)
		createDirHelper(t, dir, subDir1)

		subDir2 := "testSubDir"
		createDirHelper(t, filepath.Join(dir, subDir1), subDir2)

		fileName := subFile1 + ".json"
		f, err := ft.MoveFileToExistingDir(fileName, "testDir1/testSubDir")
		if err != nil {
			t.Errorf("got error %v, but did not want one", err)
		}

		if fileName != f {
			t.Errorf("wrong file name returned, got %s, want %s", fileName, f)
		}
		assertFileExistence(t, dir, subDir1, subDir2, fileName)
	})

	t.Run("move a file from a directory to the lab root", func(t *testing.T) {
		subDir1 := "testDir1"
		subFile1 := "testFile1"
		fileName := "testMoveToRoot.json"
		oldpath :=  subDir1+"/"+fileName
		dir, ft := createTempDir(t, "testMoveToRoot", subFile1)
		defer os.RemoveAll(dir)
		createDirHelper(t, dir, subDir1)
		createFileHelper(t, dir, oldpath)

		f, err := ft.MoveFileToExistingDir(oldpath, "/")
		if err != nil {
			t.Errorf("got error %v, but did not want one", err)
		}

		if fileName != f {
			t.Errorf("wrong file name returned, got %s, want %s", fileName, f)
		}
		assertFileExistence(t, dir)
	})

	t.Run("move a file from a directory to another but not in the lab root", func(t *testing.T) {
		subDir1 := "testDir1"
		subDir2 := "testSubDir"
		subFile1 := "testFile1"
		fileName := "testMoveToRoot.json"
		oldPath := subDir1 + "/" + fileName
		newPath := subDir1 + "/" + subDir2
		dir, ft := createTempDir(t, "testMoveToDir", subFile1)
		createDirHelper(t, dir, subDir1)
		createFileHelper(t, dir, oldPath)
		createDirHelper(t, filepath.Join(dir, subDir1), subDir2)
		defer os.RemoveAll(dir)

		f, err := ft.MoveFileToExistingDir(oldPath, newPath)
		if err != nil {
			t.Errorf("got error %v, but did not want one", err)
		}

		if fileName != f {
			t.Errorf("wrong file name returned, got %s, want %s", fileName, f)
		}

		assertFileExistence(t, dir, subDir1, subDir2)
	})

	t.Run("get an error when paths are stricty identical", func(t *testing.T) {
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testMoveToDir", subFile1)
		defer os.RemoveAll(dir)

		p := "/"
		want := ErrEqualOldAndNewPath
		_, err := ft.MoveFileToExistingDir(p, p)
		assertError(t, err, want)
	})

	t.Run("get an error when we want to move a file from the lab root to the same location", func(t *testing.T) {
		subFile1 := "testFile1"
		dir, ft := createTempDir(t, "testMoveToDir", subFile1)
		defer os.RemoveAll(dir)

		f := subFile1 + ".json"
		want := ErrEqualOldAndNewPath
		_, err := ft.MoveFileToExistingDir(f, "/")
		assertError(t, err, want)
	})

	t.Run("get an error when we want move a file from a directory, that is NOT the lab root, to the same directory", func(t *testing.T) {
		subDir1 := "testDir1"
		subFile1 := "testFile1"
		fileName := "testMoveToRoot.json"
		oldPath := subDir1 + "/" + fileName
		newPath := subDir1
		dir, ft := createTempDir(t, "testMoveToDir", subFile1)
		createDirHelper(t, dir, subDir1)
		createFileHelper(t, dir, oldPath)

		want := ErrEqualOldAndNewPath
		_, err := ft.MoveFileToExistingDir(oldPath, newPath)
		assertError(t, err, want)
	})
}

// Function used in tests
// Compare 2 files content. The goal is to fail asap
// Read files by chunks defined by the chunkSize variable
func fileContentCompare(path, file1, file2 string) error {
	const chunkSize = 2000
	path1 := filepath.Join(path, file1)
	path2 := filepath.Join(path, file2)

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