package filetree

import "testing"

func SameTest(t *testing.T) {
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
	ft := FileTreeExplorer{
		FileTree: tree,
	}

	want := true
	result := ft.Same(&tree)
	if want != result {
		t.Fatal("L'arbre n'est pas égal à lui même")
	}
}
