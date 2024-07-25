package filetree

import (
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
