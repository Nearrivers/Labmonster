package filetree

import (
	"os"
	"reflect"
	"testing"
)

func TestSaveAndOpenFile(t *testing.T) {
	t.Run("Saving a graph and opening it", func(t *testing.T) {
		want := getNewTestGraph()
		fileName := "saveopentest.json"
		ft, dir := getNewFileTreeExplorer()
		defer os.RemoveAll(dir)

		err := ft.SaveFile(fileName, want)
		if err != nil {
			t.Fatalf("got an error while saving the file but didn't want one: %v", err)
		}

		got, openErr := ft.OpenFile(fileName)
		if openErr != nil {
			t.Fatalf("got an error while opening the file but didn't want one: %v", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}