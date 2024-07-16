package walker

import (
	"flow-poc/backend/filetree"
	"testing"
)

func getTestTree() filetree.Node {
	return filetree.Node{
		Name: "Lab",
		Type: filetree.DIR,
		Files: []*filetree.Node{
			{
				Name: "Training",
				Type: filetree.DIR,
				Files: []*filetree.Node{
					{
						Name:  "Test.txt",
						Type:  filetree.FILE,
						Files: []*filetree.Node{},
					},
					{
						Name: "sous-lab",
						Type: filetree.DIR,
						Files: []*filetree.Node{
							{
								Name:  "Test2.txt",
								Type:  filetree.FILE,
								Files: []*filetree.Node{},
							},
						},
					},
				},
			},
			{
				Name: "Foo",
				Type: filetree.DIR,
				Files: []*filetree.Node{
					{
						Name:  "Bar.txt",
						Type:  filetree.FILE,
						Files: []*filetree.Node{},
					},
				},
			},
		},
	}
}
func TestSame(t *testing.T) {
	tree := getTestTree()

	t.Run("Tree compared to himself", func(t *testing.T) {
		otherTree := getTestTree()
		want := true
		got := Same(&tree, &otherTree)
		if want != got {
			t.Error("The three was compared with an identical tree and the function still found it was different")
		}
	})

	t.Run("Compared to a different tree", func(t *testing.T) {
		otherTree := filetree.Node{
			Name: "Lab",
			Type: filetree.DIR,
			Files: []*filetree.Node{
				{
					Name: "Training",
					Type: filetree.DIR,
					Files: []*filetree.Node{
						{
							Name:  "Test.txt",
							Type:  filetree.FILE,
							Files: []*filetree.Node{},
						},
					},
				},
			},
		}
		want := false
		got := Same(&tree, &otherTree)
		if want != got {
			t.Error("Trees aren't equal but function found otherwise")
		}
	})
}
