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
		newNode := InsertNode(false, &testNode, wantedName)
		if newNode.Name != wantedName {
			t.Errorf("got %s want %s", wantedName, newNode.Name)
		}

		if newNode.Type != FILE {
			t.Errorf("")
		}
	})
}

func TestRemoveNode(t *testing.T) {
	t.Run("Node deletion", func(t *testing.T) {
		nodeName := "Node to delete"
		testNode := Node{
			Name: "Lab",
			Type: DIR,
			Files: []*Node{
				{
					Name:  nodeName,
					Type:  FILE,
					Files: []*Node{},
				},
			},
		}

		want := Node{
			Name:  "Lab",
			Type:  DIR,
			Files: []*Node{},
		}

		got, err := RemoveNode(&testNode, nodeName)
		if err != nil {
			t.Errorf("An error occured %q", err)
		}

		result := Same(&want, got)
		if !result {
			t.Errorf("%+v and %+v should be the same", got, want)
		}
	})

	t.Run("Deleting a node that does not exists", func(t *testing.T) {
		nodeName := "non-existant node"
		testNode := Node{
			Name: "Lab",
			Type: DIR,
			Files: []*Node{
				{
					Name:  "t1",
					Type:  FILE,
					Files: []*Node{},
				},
			},
		}

		_, got := RemoveNode(&testNode, nodeName)
		if got == nil {
			t.Error("An error should've occured")
		}

		want := ErrNodeNotFound
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestRemoveIndex(t *testing.T) {
	t.Run("Existing index deletion", func(t *testing.T) {
		nodes := []*Node{
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

		got, err := removeIndex(nodes, 1)
		if err != nil {
			t.Errorf("An error occured %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %+v", got, want)
		}
	})

	t.Run("Out of bounds index deletion", func(t *testing.T) {
		nodes := []*Node{
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
		}

		_, got := removeIndex(nodes, 99)
		if got == nil {
			t.Error("An occured should've occured")
		}

		want := ErrIndexOutOfBounds
		if got != ErrIndexOutOfBounds {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
