package filetree_test

import (
	"flow-poc/backend/config"
	"flow-poc/backend/filetree"
	"os"
	"testing"
)

func createTempDir(t testing.TB, dirName string) (string, *filetree.FileTree) {
	t.Helper()

	dir, err := os.MkdirTemp("", dirName)
	if err != nil {
		t.Fatalf("an error occured while creating temporary directory: %v", err)
	}

	ft := filetree.NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	})

	return dir, ft
}

func openPngImageFile(t testing.TB) string {
	b, err := os.ReadFile("./testFiles/pngImage.txt")
	if err != nil {
		t.Fatalf("couldn't find pngImage file: %v", err)
	}

	return string(b)
}

func TestSaveMedia(t *testing.T) {
	t.Run("save png image", func(t *testing.T) {
		dir, ft := createTempDir(t, "savePng")
		defer os.RemoveAll(dir)
		b := openPngImageFile(t)

		path, err := ft.SaveMedia("", "image/png", b)
		if err != nil {
			t.Fatalf("got an unexpected error: %v", err)
		}

		if _, err = os.Stat(path); err != nil {
			t.Errorf("didn't find file but should have")
		}
	})
}
