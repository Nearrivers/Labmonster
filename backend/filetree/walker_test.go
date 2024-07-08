package filetree

import "testing"

func TestSame(t *testing.T) {
	tree := Node{
		Name: "Lab",
		Type: DIR,
		Files: []*Node{
			{
				Name: "Training",
				Type: DIR,
				Files: []*Node{
					{
						Name:  "Test.txt",
						Type:  FILE,
						Files: []*Node{},
					},
					{
						Name: "sous-lab",
						Type: DIR,
						Files: []*Node{
							{
								Name:  "Test2.txt",
								Type:  FILE,
								Files: []*Node{},
							},
						},
					},
				},
			},
			{
				Name: "Foo",
				Type: DIR,
				Files: []*Node{
					{
						Name:  "Bar.txt",
						Type:  FILE,
						Files: []*Node{},
					},
				},
			},
		},
	}

	t.Run("Tree compared to himself", func(t *testing.T) {
		want := true
		got := Same(&tree, &tree)
		if want != got {
			t.Error("The three was compared with himself and the function still found it was different")
		}
	})

	t.Run("Compared to a different tree", func(t *testing.T) {
		otherTree := Node{
			Name: "Lab",
			Type: DIR,
			Files: []*Node{
				{
					Name: "Training",
					Type: DIR,
					Files: []*Node{
						{
							Name:  "Test.txt",
							Type:  FILE,
							Files: []*Node{},
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
