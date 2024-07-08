package filetree

import "testing"

func TestInsertNode(t *testing.T) {
	t.Run("File insertion", func(t *testing.T) {
		testNode := Node{
			Name:  "Lab",
			Type:  DIR,
			Files: []*Node{},
		}

		wantedName := "nodetest"
		newNode := InsertNode(false, &testNode, wantedName)
		if newNode.Name == wantedName {
			t.Errorf("got %s want %s", wantedName, newNode.Name)
		}

		if newNode.Type != FILE {
			t.Errorf("")
		}
	})
}
