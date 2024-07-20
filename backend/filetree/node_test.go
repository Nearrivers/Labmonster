package filetree

import (
	"reflect"
	"testing"
)

func TestInsertNode(t *testing.T) {
	t.Run("Node insertion", func(t *testing.T) {
		testNode := Node{
			Name:  "Lab",
			Type:  DIR,
			Files: []*Node{},
		}

		wantedName := "nodetest"
		newNode := testNode.InsertNode(false, wantedName)
		if newNode.Name != wantedName {
			t.Errorf("got %s want %s", wantedName, newNode.Name)
		}

		if newNode.Type != FILE {
			t.Errorf("")
		}
	})
}

func TestRemoveIndex(t *testing.T) {
	t.Run("Existing index deletion", func(t *testing.T) {
		n := &Node{
			Name: "test",
			Type: DIR,
			Files: []*Node{
				{
					Name: "t1",
					Type: FILE,
				},
				{
					Name: "t2",
					Type: FILE,
				},
				{
					Name: "t3",
					Type: FILE,
				},
			},
		}

		want := []*Node{
			{
				Name: "t1",
				Type: FILE,
			},
			{
				Name: "t3",
				Type: FILE,
			},
		}

		err := n.removeIndex(1)
		if err != nil {
			t.Errorf("An error occured %v", err)
		}

		if !reflect.DeepEqual(n.Files, want) {
			t.Errorf("got %#v want %+v", n.Files, want)
		}
	})

	t.Run("Out of bounds index deletion", func(t *testing.T) {
		n := &Node{
			Name: "test",
			Type: DIR,
			Files: []*Node{
				{
					Name: "t1",
					Type: FILE,
				},
				{
					Name: "t2",
					Type: FILE,
				},
				{
					Name: "t3",
					Type: FILE,
				},
			},
		}

		got := n.removeIndex(99)
		if got == nil {
			t.Error("An occured should've occured")
		}

		want := ErrIndexOutOfBounds
		if got != ErrIndexOutOfBounds {
			t.Errorf("want %v got %v", want, got)
		}
	})
}

func TestSortNodes(t *testing.T) {}
