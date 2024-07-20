package filetree

import (
	"os"
	"path/filepath"
	"testing"

	"flow-poc/backend/config"
)

func createDirAndFile(t testing.TB, tempDirPath, dirName string) {
	t.Helper()

	err := os.Mkdir(filepath.Join(tempDirPath, dirName), 0750)
	if err != nil {
		t.Fatalf("couldn't create the first subdirectory: %v", err)
	}
}

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
		createDirAndFile(t, dir, subDir1)

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
		createDirAndFile(t, dir, subDir1)

		ft.FileTree.Files = append(ft.FileTree.Files, &Node{
			Name:  subDir1,
			Type:  DIR,
			Files: make([]*Node, 0),
		})

		subDir2 := "testSubDir"
		createDirAndFile(t, filepath.Join(dir, subDir1), subDir2)

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

func TestFindAndDeleteNode(t *testing.T) {
	t.Run("find and delete node at second depth", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())

		ft.FileTree.Files = []*Node{
			{
				Name: "testDir",
				Type: DIR,
				Files: []*Node{
					{
						Name: "nodeToDelete.json",
						Type: FILE,
					},
				},
			},
		}

		err := ft.findAndDeleteNode("nodeToDelete.json", ft.FileTree.Files[0].Files)
		if err != nil {
			t.Errorf("got an error but should not have: %v", err)
		}

		n := ft.FileTree.Files[0]
		if len(n.Files) != 0 {
			t.Error("the node was not deleted")
		}
	})

	t.Run("find and delete node at first depth", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())

		ft.FileTree.Files = []*Node{
			{
				Name:  "testDir",
				Type:  DIR,
				Files: []*Node{},
			},
		}

		err := ft.findAndDeleteNode("testDir", ft.FileTree.Files)
		if err != nil {
			t.Errorf("got an error but should not have: %v", err)
		}

		n := ft.FileTree
		if len(n.Files) != 0 {
			t.Error("the node was not deleted")
		}
	})

	t.Run("trying to find a node that does not exists", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())

		ft.FileTree.Files = []*Node{
			{
				Name:  "testDir",
				Type:  DIR,
				Files: []*Node{},
			},
		}

		err := ft.findAndDeleteNode("testDir23", ft.FileTree.Files)
		if err == nil {
			t.Error("didn't get an error but should have")
		}
	})

	t.Run("sending empty path", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())

		ft.FileTree.Files = []*Node{
			{
				Name:  "testDir",
				Type:  DIR,
				Files: []*Node{},
			},
		}

		err := ft.findAndDeleteNode("", ft.FileTree.Files)
		if err == ErrPathMissing {
			t.Errorf("got %v, want %v", err, ErrPathMissing)
		}
	})
}

func TestRenameNode(t *testing.T) {
	t.Run("rename a node at second depth", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())
		newNodeName := "newNodeName"
		oldNodeName := "nodeToRename"

		ft.FileTree.Files = []*Node{
			{
				Name: "testDir",
				Type: DIR,
				Files: []*Node{
					{
						Name: oldNodeName + ".json",
						Type: FILE,
					},
				},
			},
		}

		err := ft.RenameNode("testDir", "nodeToRename", newNodeName)
		if err != nil {
			t.Errorf("got an error but didn't expect one: %v", err)
		}

		got := ft.FileTree.Files[0].Files[0].Name

		if got != newNodeName+".json" {
			t.Errorf("the name of the node didn't change, got %s, want %s", got, newNodeName+".json")
		}
	})

	t.Run("rename a node at first depth", func(t *testing.T) {
		ft := NewFileTree(config.NewAppConfig())
		newNodeName := "newNodeName"
		oldNodeName := "nodeToRename"

		ft.FileTree.Files = []*Node{
			{
				Name:  "testDir",
				Type:  DIR,
				Files: []*Node{},
			},

			{
				Name: oldNodeName + ".json",
				Type: FILE,
			},
		}

		err := ft.RenameNode("", "nodeToRename", newNodeName)
		if err != nil {
			t.Fatalf("got an error but didn't expect one: %v", err)
		}

		got := ft.FileTree.Files[0].Files[0].Name

		if got != newNodeName+".json" {
			t.Errorf("the name of the node didn't change, got %s, want %s", got, newNodeName+".json")
		}
	})
}
