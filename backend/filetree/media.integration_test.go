package filetree_test

import (
	"flow-poc/backend/config"
	"flow-poc/backend/filetree"
	"os"
	"testing"
)

func saveMedia(t testing.TB, mime string) (dir, path string, ft *filetree.FileTree) {
	t.Helper()

	dir, ft = createTempDir(t, "savePng")
	s := openPngImageFile(t)

	path, err := ft.SaveMedia("", mime, s)
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}

	if _, err = os.Stat(path); err != nil {
		t.Fatalf("didn't find file but should have")
	}

	return
}

func TestSaveAndOpenMediaConc(t *testing.T) {
	t.Run("save png image and open it concurrently", func(t *testing.T) {
		dir, path, ft := saveMedia(t, "image/png")
		defer os.RemoveAll(dir)

		got, err := ft.OpenMediaConc(path)
		if err != nil {
			t.Errorf("got an error but didn't expect one: %s", err)
		}

		want, err := ft.OpenMedia(path)
		if err != nil {
			t.Fatalf("got an error but didn't expect one: %s", err)
		}

		if got[:50] != want[:50] {
			t.Errorf("got %s, want %s", got[:50], want[:50])
		}
	})
}

func BenchmarkOpenMedia(b *testing.B) {
	ft := filetree.NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{LabPath: ""},
	})

	b.Run("benchmark synchronous open media", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ft.OpenMedia("./testFiles/Wall jab combo.mp4")
		}
	})

	b.Run("benchmark concurrent open media", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ft.OpenMediaConc("./testFiles/Wall jab combo.mp4")
		}
	})
}
