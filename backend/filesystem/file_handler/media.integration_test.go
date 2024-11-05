package file_handler_test

import (
	"flow-poc/backend/filesystem/file_handler"
	"os"
	"testing"
)

func saveMedia(t testing.TB, mime string) (dir, path string, ft *file_handler.FileHandler) {
	t.Helper()

	dir, ft = createTempDir(t, "savePng")
	s := openPngImageFile(t)

	path, err := ft.SaveMedia("", "", mime, s)
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}

	if _, err = os.Stat(path); err != nil {
		t.Fatalf("didn't find file but should have")
	}

	return
}
