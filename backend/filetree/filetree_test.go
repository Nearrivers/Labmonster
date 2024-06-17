package filetree

import (
	"testing"

	"flow-poc/backend/config"
)

func TestGetFileTree(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	want := Node{
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
								Files: []*Node{},
								Type:  FILE,
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

	ft2 := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})
	ft2.FileTree = want

	_, err := ft.GetFileTree()
	if err != nil {
		t.Fatal(err)
	}

	if !ft.Same(&want) {
		ft.PrintTree()
		ft.Logger.Debug("-----------")
		ft2.PrintTree()
		t.Fatal("Les 2 arbres ne sont pas égaux")
	}
}