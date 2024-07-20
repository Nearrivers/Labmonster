package filetree

import (
	"errors"
	"flow-poc/backend/config"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getNewFileTreeExplorer() *FileTreeExplorer {
	return NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test",
		},
	})
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
		ft := getNewFileTreeExplorer()

		testFileName := "happyPath test"
		defer ft.DeleteFile(testFileName)

		got, err := ft.CreateNewFileAtRoot(testFileName)
		if err != nil {
			t.Fatalf("An error occured while creating the file: %v", err.Error())
		}

		_, err = os.Stat(filepath.Join(ft.GetLabPath(), got))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Fatal("The file was not found after its supposed creation")
			}
			t.Fatalf("An error occured while using os.Stat: %v", err.Error())
		}
	})

	t.Run("Creating multiple files in a row", func(t *testing.T) {
		ft := getNewFileTreeExplorer()

		testFileName := "Multiple test"

		cpt := 0
		for cpt < 3 {
			newFileName, err := ft.CreateNewFileAtRoot(testFileName)
			if err != nil {
				t.Errorf("An error occured while creating the file: %s", err.Error())
			}
			defer ft.DeleteFile(newFileName)

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

	t.Run("Checking node presence inside in-memory tree", func(t *testing.T) {
		ft := getNewFileTreeExplorer()

		fileName := "in-memory test"
		_, err := ft.CreateNewFileAtRoot(fileName)
		if err != nil {
			t.Errorf("An error occured while creating the file: %v", err)
		}

		defer ft.DeleteFile(fileName)

		_, _, err = searchFileOrDir(fileName+".json", ft.FileTree.Files)
		if err != nil {
			t.Error("The node wasn't found inside the in-memory tree")
		}
	})
}

func TestSearchFile(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Error("An error should have occured")
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Searching for existing node", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		n := &ft.FileTree
		n.InsertNode(false, "Test 1")
		n.InsertNode(false, "Test 2")
		n.InsertNode(false, "Test 3")
		n.InsertNode(false, "Test 4")

		want := "Test 3"
		wantedIndex := 2
		file, index, err := searchFileOrDir(want, ft.FileTree.Files)
		if err != nil {
			t.Error(err.Error())
		}

		got := file.Name
		gotIndex := index

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if gotIndex != wantedIndex {
			t.Errorf("Indexes are different: got %d, want %d", gotIndex, wantedIndex)
		}
	})

	t.Run("Searching in an empty level", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		want := ErrNoFileInThisLevel
		_, _, got := searchFileOrDir("Test 3", ft.FileTree.Files)
		assertError(t, got, want)
	})

	t.Run("Searching for a node that doesn't exist", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		n := &ft.FileTree
		n.InsertNode(false, "Test 1")
		n.InsertNode(false, "Test 2")
		n.InsertNode(false, "Test 3")
		n.InsertNode(false, "Test 4")

		want := ErrNodeNotFound
		_, _, got := searchFileOrDir("Test 5", ft.FileTree.Files)
		assertError(t, got, want)
	})
}

func TestRenameFile(t *testing.T) {
	t.Run("Existing file rename at first level", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
		oldName := "test rename"
		newName := "test rename 2"

		createFileBeforeTest(t, ft, oldName)
		defer ft.DeleteFile(newName)

		err := ft.RenameFile("", oldName+".json", newName)
		if err != nil {
			t.Errorf("got an error %v", err)
		}
	})

	t.Run("Non existant file rename at first level", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
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
		ft := getNewFileTreeExplorer()
		fileName := "test delete"

		createFileBeforeTest(t, ft, fileName)

		err := ft.DeleteFile(fileName)
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
		ft := getNewFileTreeExplorer()
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

// TODO: Une fois la sauvegarde des fichiers implémentée, tester la duplication
// avec des fichiers modifiés afin de voir pour de bon si le contenu des fichiers
// est bel et bien dupliqué
func TestDuplicateFile(t *testing.T) {
	t.Run("File duplication at first level", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
		fileName := "duplication test"
		createFileBeforeTest(t, ft, fileName)
		defer ft.DeleteFile(fileName)

		duplicatedFile, err := ft.DuplicateFile("", fileName)
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}
		defer ft.DeleteFile(duplicatedFile)

		if err = fileContentCompare(ft, fileName+".json", duplicatedFile); err != nil {
			if err == ErrFileAreDifferent {
				t.Fatalf("the two files are different: %v", err)
			}

			t.Errorf("an error occured before comparing the two files: %v", err)
		}
	})

	t.Run("Running multiple file duplications in a row", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
		fileName := "multiple duplication test"
		createFileBeforeTest(t, ft, fileName)
		defer ft.DeleteFile(fileName)

		duplicatedFile1, err := ft.DuplicateFile("", fileName)
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}

		duplicatedFile2, err := ft.DuplicateFile("", fileName)
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}

		ft.DeleteFile(duplicatedFile1)
		ft.DeleteFile(duplicatedFile2)
	})

	t.Run("Trying to duplicate a file that doesn't exists", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
		fileName := "fake duplication test"

		_, err := ft.DuplicateFile("", fileName)
		if err == nil {
			t.Fatal("didn't get an error, but wanted one")
		}
	})
}
