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

	want := true
	result := Same(&tree, &tree)
	if want != result {
		t.Fatal("Un arbre vient d'être comparé à lui-même et Same() à trouvé qu'il été quand même différent")
	}
}
