package recentfiles

import (
	"bufio"
	"flow-poc/backend/config"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func initRecentlyOpened(t testing.TB, max int) (*RecentlyOpened, string) {
	t.Helper()
	dir, err := os.MkdirTemp("", "recentTests")
	if err != nil {
		t.Fatalf("an error occured while creating temp dir for tests: %v", err)
	}

	err = os.Mkdir(filepath.Join(dir, ".labmonster"), os.ModeAppend)
	if err != nil {
		t.Fatalf("an error occured while creating .labmonster dir for tests: %v", err)
	}

	return NewRecentlyOpened(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: dir,
		},
	}, max), dir
}

func assertRecentSaved(t testing.TB, tempDirPath string) {
	t.Helper()

	_, err := os.Stat(filepath.Join(tempDirPath, ".labmonster", recentlyOpenedFilename))
	if err != nil {
		t.Fatalf("couldn't find recently opended file: %v", err)
	}
}

func assertContentMatchWithSaved(t testing.TB, tempDirPath string, r *RecentlyOpened) {
	t.Helper()

	f, err := os.Open(filepath.Join(tempDirPath, ".labmonster", recentlyOpenedFilename))
	if err != nil {
		t.Fatalf("couldn't open recently opended file: %v", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	cpt := 0
	for scanner.Scan() {
		got := scanner.Text()
		want := r.FilePaths[cpt]

		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
		cpt++
	}

	if err = scanner.Err(); err != nil {
		t.Fatalf("couldn't read recently opended file: %v", err)
	}
}

func TestSaveRecent(t *testing.T) {
	r, dir := initRecentlyOpened(t, 10)
	defer os.RemoveAll(dir)

	tests := []struct {
		name        string
		filesToSave []string
	}{
		{
			name:        "Save one recent file",
			filesToSave: []string{"testRecent.json"},
		},
		{
			name:        "Save multiple files and one not at lab root",
			filesToSave: []string{"testRecent.json", "/Foo/testRecentFoo.json"},
		},
		{
			name:        "Save multiple non json files",
			filesToSave: []string{"/Bar/Foo/Fizz/testRecentDeep.json", "pasUnJSON.png", "/Bar/text.txt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.FilePaths = tt.filesToSave
			err := r.SaveRecentlyOpended()
			if err != nil {
				t.Errorf("got an error but didn't want one: %v", err)
			}

			assertRecentSaved(t, dir)
			assertContentMatchWithSaved(t, dir, r)
		})
	}
}

func TestAddRecent(t *testing.T) {
	t.Run("add recent file", func(t *testing.T) {
		r, dir := initRecentlyOpened(t, 1)
		defer os.RemoveAll(dir)
		file := "testAddRecent.json"
		p := filepath.Join(dir, file)

		r.AddRecentFile(p)
		if !slices.Contains(r.FilePaths, p) {
			t.Errorf("did not insert the file in the array")
		}
	})

	t.Run("add the same file twice in a row", func(t *testing.T) {
		r, dir := initRecentlyOpened(t, 2)
		defer os.RemoveAll(dir)
		file := "testAddRecentTwice.json"
		p := filepath.Join(dir, file)

		r.AddRecentFile(p)
		r.AddRecentFile(p)

		if !slices.Contains(r.FilePaths, p) {
			t.Fatal("did not insert the file in the array")
		}

		if len(r.FilePaths) == 2 {
			t.Error("the same path cannot be inserted twice")
		}
	})

	t.Run("add more recent files than expected", func(t *testing.T) {
		r, dir := initRecentlyOpened(t, 2)
		defer os.RemoveAll(dir)
		f1 := "testAddRecentTwice.json"
		f2 := "testAddTooMuch2.json"
		f3 := "testAddTooMuch3.json"
		p1 := filepath.Join(dir, f1)
		p2 := filepath.Join(dir, f2)
		p3 := filepath.Join(dir, f3)

		want := []string{p3, p2}

		r.AddRecentFile(p1)
		r.AddRecentFile(p2)
		r.AddRecentFile(p3)

		if slices.Compare(r.FilePaths, want) != 0 {
			t.Errorf("got %v, want %v", r.FilePaths, want)
		}
	})
}
