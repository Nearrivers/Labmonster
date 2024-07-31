package filetree

import (
	"errors"
	"flow-poc/backend/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

		testFileName := "happyPath test"
		defer ft.DeleteFile(testFileName)

		got, err := ft.CreateNewFileAtRoot(testFileName)
		if err != nil {
			t.Fatalf("An error occured while creating the file: %v", err.Error())
		}

		_, err = os.Stat(filepath.Join(ft.GetLabPath(), got.Name+".json"))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Fatalf("The file %s was not found after its supposed creation", got.Name)
			}
			t.Fatalf("An error occured while using os.Stat: %v", err.Error())
		}
	})

	t.Run("Creating multiple files in a row", func(t *testing.T) {
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		testFileName := "Multiple test"

		cpt := 0
		for cpt < 3 {
			newFile, err := ft.CreateNewFileAtRoot(testFileName)
			if err != nil {
				t.Errorf("An error occured while creating the file: %s", err.Error())
			}
			defer ft.DeleteFile(newFile.Name)

			if cpt > 0 {
				_, err = os.Stat(filepath.Join(ft.GetLabPath(), fmt.Sprintf(testFileName+" %d.json", cpt)))
			} else {
				_, err = os.Stat(filepath.Join(ft.GetLabPath(), testFileName+".json"))
			}

			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					t.Fatal("The file was not found after its supposed creation")
				}
				t.Errorf("An error occured while using os.Stat: %v", err)
			}

			cpt++
		}
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

		fileName := "test delete.json"

		createFileBeforeTest(t, ft, fileName)

		err := ft.DeleteFile(fileName)
		if err != nil {
			t.Errorf("An error occured while deleting the file: %v", err)
		}

		want := false
		got := doesFileExist(filepath.Join(ft.GetLabPath(), fileName))

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

	if err := ft.fileContentCompare(file1, file2); err != nil {
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