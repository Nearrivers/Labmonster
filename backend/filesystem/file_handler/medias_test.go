package file_handler_test

import (
	"flow-poc/backend/config"
	"flow-poc/backend/filesystem/file_handler"
	"os"
	"path/filepath"
	"testing"
)

const pngFileName = "testImage.png"

func createTempDir(t testing.TB, dirName string) (string, *file_handler.FileHandler) {
	t.Helper()

	dir, err := os.MkdirTemp("", dirName)
	if err != nil {
		t.Fatalf("an error occured while creating temporary directory: %v", err)
	}

	fh := file_handler.NewFileHandler(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	})

	b, err := os.ReadFile("./testFiles/pngImage.txt")
	if err != nil {
		t.Fatalf("couldn't find pngImage file: %v", err)
	}

	f, cErr := os.Create(filepath.Join(fh.GetLabPath(), pngFileName))
	if cErr != nil {
		t.Fatalf("couldn't create test image file: %v", cErr)
	}
	defer f.Close()

	_, wErr := f.Write(b)
	if wErr != nil {
		t.Fatalf("couldn't write in test image file: %v", wErr)
	}

	return dir, fh
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
		s := openPngImageFile(t)

		path, err := ft.SaveMedia("", "", "image/png", s)
		if err != nil {
			t.Fatalf("got an unexpected error: %v", err)
		}

		if _, err = os.Stat(path); err != nil {
			t.Errorf("didn't find file but should have")
		}
	})

	// TODO: Tester les autres formats de fichiers
}

func TestOpenMedia(t *testing.T) {
	t.Run("opening a png image with absolute path", func(t *testing.T) {
		dir, ft := createTempDir(t, "openPng")
		defer os.RemoveAll(dir)
		p := filepath.Join(dir, pngFileName)

		_, err := ft.OpenMedia(p)
		if err != nil {
			t.Errorf("couldn't open media: %v", err)
		}
	})

	t.Run("opening a png image with path relative the lab's root", func(t *testing.T) {
		dir, ft := createTempDir(t, "openPng")
		defer os.RemoveAll(dir)

		_, err := ft.OpenMedia(pngFileName)
		if err != nil {
			t.Errorf("couldn't open media: %v", err)
		}
	})
}
