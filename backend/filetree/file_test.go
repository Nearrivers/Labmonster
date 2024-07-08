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

func TestCreateNewFile(t *testing.T) {
	t.Run("Happy path test", func(t *testing.T) {
		ft := getNewFileTreeExplorer()

		testFileName := "happyPath test"

		want := testFileName + ".json"
		defer ft.DeleteFile(want)

		got, err := ft.CreateNewFile(testFileName)
		if err != nil {
			t.Errorf("An error occured while creating the file: %v", err.Error())
		}

		_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, want))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Error("The file was not found after its supposed creation")
			}
			t.Errorf("An error occured while using os.Stat: %v", err.Error())
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Creating multiple files in a row", func(t *testing.T) {
		ft := getNewFileTreeExplorer()

		testFileName := "Multiple test"
		defer ft.DeleteFile(testFileName + ".json")
		defer ft.DeleteFile(testFileName + " 1.json")
		defer ft.DeleteFile(testFileName + " 2.json")

		cpt := 0
		for cpt < 3 {
			_, err := ft.CreateNewFile(testFileName)
			if err != nil {
				t.Errorf("An error occured while creating the file: %s", err.Error())
			}

			if cpt > 0 {
				_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, fmt.Sprintf(testFileName+" %d.json", cpt)))
			} else {
				_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, testFileName+".json"))
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

		InsertNode(false, &ft.FileTree, "Test 1")
		InsertNode(false, &ft.FileTree, "Test 2")
		InsertNode(false, &ft.FileTree, "Test 3")
		InsertNode(false, &ft.FileTree, "Test 4")

		want := "Test 3"
		wantedIndex := 2
		file, index, err := searchFile(want, ft.FileTree.Files)
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
		_, _, got := searchFile("Test 3", ft.FileTree.Files)
		assertError(t, got, want)
	})

	t.Run("Searching for a node that doesn't exist", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		InsertNode(false, &ft.FileTree, "Test 1")
		InsertNode(false, &ft.FileTree, "Test 2")
		InsertNode(false, &ft.FileTree, "Test 3")
		InsertNode(false, &ft.FileTree, "Test 4")

		want := ErrNodeNotFound
		_, _, got := searchFile("Test 5", ft.FileTree.Files)
		assertError(t, got, want)
	})
}

func TestDeleteFile(t *testing.T) {
	t.Run("Existing file deletion at first level", func(t *testing.T) {
		ft := getNewFileTreeExplorer()
		fileName := "test delete"

		_, err := ft.CreateNewFile(fileName)
		if err != nil {
			t.Error("An error occured while creating the file but should've have")
		}

		err = ft.DeleteFile(fileName + ".json")
		if err != nil {
			t.Errorf("An error occured while deleting the file: %v", err)
		}

		want := false
		got := doesFileExist(filepath.Join(ft.Cfg.ConfigFile.LabPath, fileName+".json"))

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
