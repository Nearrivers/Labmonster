package filetree

// import (
// 	"testing"

// 	"flow-poc/backend/config"
// )

// Test très peu fiable, préparer un meilleur environnement avec des dossiers temporaires afin de créer un vrai test
// ou carrément partir du pwd (project working directory)
// func TestGetFileTree(t *testing.T) {
// 	ft := NewFileTree(&config.AppConfig{
// 		ConfigFile: config.ConfigFile{
// 			LabPath: "D:\\Projets\\test\\Lab",
// 		},
// 	})

// 	want := Node{
// 		Name: "Lab",
// 		Type: DIR,
// 		Files: []*Node{
// 			{
// 				Name: "Training",
// 				Type: DIR,
// 				Files: []*Node{
// 					{
// 						Name:  "Test.txt",
// 						Type:  FILE,
// 						Files: []*Node{},
// 					},
// 					{
// 						Name: "sous-lab",
// 						Type: DIR,
// 						Files: []*Node{
// 							{
// 								Name:  "Test2.txt",
// 								Files: []*Node{},
// 								Type:  FILE,
// 							},
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Name: "Foo",
// 				Type: DIR,
// 				Files: []*Node{
// 					{
// 						Name:  "Bar.txt",
// 						Type:  FILE,
// 						Files: []*Node{},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	ft2 := NewFileTree(&config.AppConfig{
// 		ConfigFile: config.ConfigFile{
// 			LabPath: "D:\\Projets\\test\\Lab",
// 		},
// 	})
// 	ft2.FileTree = want

// 	_, err := ft.GetTheWholeTree()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	PrintTree(ft.FileTree)
// 	ft.Logger.Debug("-----------")
// 	PrintTree(ft2.FileTree)

// 	if !Same(&ft.FileTree, &want) {
// 		t.Fatal("Les 2 arbres ne sont pas égaux")
// 	}
// }
