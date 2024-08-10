package recentfiles

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func simulateAppShutdown(r *RecentlyOpened) {
	r.FilePaths = []string{}
}

func TestAddSaveLoad(t *testing.T) {
	t.Run("Add a recently opended file, save it and load it back", func(t *testing.T) {
		r, dir := initRecentlyOpened(t, 5)
		defer os.RemoveAll(dir)
		f1 := "/foo/testAddSaveLoad.json"
		f2 := "setup.mp4"
		p1 := filepath.Join(dir, f1)
		p2 := filepath.Join(dir, f2)

		r.AddRecentFile(p1)
		r.AddRecentFile(p2)
		r.SaveRecentlyOpended()
		simulateAppShutdown(r)

		err := r.LoadRecentlyOpended()
		if err != nil {
			t.Fatalf("got an error but didn't expect one: %v", err)
		}

		want := []string{p2, p1}
		if slices.Compare(r.FilePaths, want) != 0 {
			t.Errorf("got %v, want %v", r.FilePaths, want)
		}
	})
}
