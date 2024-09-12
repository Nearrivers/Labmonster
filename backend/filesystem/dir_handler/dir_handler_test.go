package dirhandler

import (
	"flow-poc/backend/config"
	"flow-poc/backend/filesystem/node"
	"flow-poc/backend/filesystem/recentfiles"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
func createTempDir(t testing.TB, testDirName string) (string, *DirHandler) {
	t.Helper()

	tempDir := os.TempDir()

	dir, err := os.MkdirTemp("", testDirName)
	if err != nil {
		t.Fatalf("an error occured while creating temporary directory: %v %s", err, tempDir)
	}

	c := &config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	}
	dh := NewDirHandler(c, recentfiles.NewRecentlyOpened(c, 5))

	return dir, dh
}

func TestGetLabDirs(t *testing.T) {
	t.Run("get every directory of the lab", func(t *testing.T) {
		subDir1 := "testDir1"
		dir, ft := createTempDir(t, "testNextLevel")
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

func assertDirExistence(t testing.TB, absPath string) {
	t.Helper()

	stat, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("the directory was not created")
		}

		t.Fatalf("couldn't get dir stats: %v", err)
	}

	if !stat.IsDir() {
		t.Fatal("the path doesn't point to a directory")
	}
}

func assertDirDoesNotExists(t testing.TB, absPath string) {
	t.Helper()

	_, err := os.Stat(absPath)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("wrong error received: %v", err)
	}
}

func TestCreateDir(t *testing.T) {
	t.Run("Creating one directory", func(t *testing.T) {
		subDir := "testDir"
		dir, dh := createTempDir(t, "testCreateDir")
		defer os.RemoveAll(dir)

		n, err := dh.CreateDirectory(subDir)
		if err != nil {
			t.Fatalf("couldn't create folder: %v", err)
		}

		if n.Type != node.DIR {
			t.Fatalf("the created node is not a directory")
		}

		if n.Name != subDir {
			t.Fatalf("wrong node name received: got %s, want %s", n.Name, subDir)
		}
	})

	t.Run("Creating multiple dirs in a row with the same name", func(t *testing.T) {
		subDir := "testDirMultiple"
		dir, dh := createTempDir(t, "testCreateDirMultiple")
		defer os.RemoveAll(dir)

		cpt := 0
		for cpt < 3 {
			_, cErr := dh.CreateDirectory(subDir)
			if cErr != nil {
				t.Fatalf("Couldn't create folder: %v", cErr)
			}

			cpt++
		}

		fp := filepath.Join(dir, subDir)
		fp1 := filepath.Join(dir, subDir+" 1")
		fp2 := filepath.Join(dir, subDir+" 2")
		assertDirExistence(t, fp)
		assertDirExistence(t, fp1)
		assertDirExistence(t, fp2)
	})

	t.Run("Creating a dir then another inside it", func(t *testing.T) {
		subDir := "testSubDir"
		subSubDir := "testSubSubDir"
		dir, dh := createTempDir(t, "testSubDir")
		defer os.RemoveAll(dir)

		_, c1Err := dh.CreateDirectory(subDir)
		if c1Err != nil {
			t.Fatalf("couldn't create first sub directory: %v", c1Err)
		}

		fp := filepath.Join(dir, subDir)
		assertDirExistence(t, fp)

		sP := filepath.Join(subDir, subSubDir)
		_, c2Err := dh.CreateDirectory(sP)
		if c2Err != nil {
			t.Fatalf("couldn't create the last sub directory: %v", c2Err)
		}

		fp2 := filepath.Join(dir, subDir, subSubDir)
		assertDirExistence(t, fp2)
	})
}

func TestRemoveDir(t *testing.T) {
	subDir := "removeDir"
	dir, dh := createTempDir(t, "testRemove")
	defer os.RemoveAll(dir)
	createDirHelper(t, dir, subDir)

	err := dh.DeleteDirectory(subDir)
	if err != nil {
		t.Errorf("couldn't delete the directory: %v", err)
	}

	fp := filepath.Join(dir, subDir)
	assertDirDoesNotExists(t, fp)
}

func TestRenameDirectory(t *testing.T) {
	subDir := "renameDir"
	dir, dh := createTempDir(t, "testRename")
	defer os.RemoveAll(dir)
	createDirHelper(t, dir, subDir)

	newSubDirName := "newNameDir"
	err := dh.RenameDirectory(subDir, newSubDirName)
	if err != nil {
		t.Errorf("couldn't rename dir: %v", err)
	}
}
