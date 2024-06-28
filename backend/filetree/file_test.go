package filetree

import (
	"errors"
	"flow-poc/backend/config"
	"os"
	"path/filepath"
	"testing"
)

func TestBasicFileCreation(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	defer ft.DeleteFile("Sans titre.json")

	err := ft.CreateNewFile()
	if err != nil {
		t.Fatalf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
	}

	_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json"))
	if err != nil {
		t.Fatalf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
	}

	if errors.Is(err, os.ErrNotExist) {
		t.Fatal("Le fichier n'a pas été créé")
	}
}

// func TestCreateMultipleFilesInARow(t *testing.T) {
// 	ft := NewFileTree(&config.AppConfig{
// 		ConfigFile: config.ConfigFile{
// 			LabPath: "D:\\Projets\\test\\Lab",
// 		},
// 	})

// 	defer ft.DeleteFile("Sans titre.json")
// 	defer ft.DeleteFile("Sans titre 1.json")
// 	defer ft.DeleteFile("Sans titre 2.json")

// 	cpt := 0
// 	for cpt < 3 {
// 		err := ft.CreateNewFile()
// 		if err != nil {
// 			t.Fatalf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
// 		}

// 		if cpt > 0 {
// 			_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, fmt.Sprintf("Sans titre %d.json", cpt)))
// 		} else {
// 			_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json"))
// 		}

// 		if err != nil {
// 			t.Fatalf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
// 		}

// 		if errors.Is(err, os.ErrNotExist) {
// 			t.Fatal("Le fichier n'a pas été créé")
// 		}

// 		cpt++
// 	}
// }

// Test du happy path
func TestSearchFile(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	InsertNode(false, &ft.FileTree, "Test 1")
	InsertNode(false, &ft.FileTree, "Test 2")
	InsertNode(false, &ft.FileTree, "Test 3")
	InsertNode(false, &ft.FileTree, "Test 4")

	file, err := searchFile("Test 3", ft.FileTree.Files)
	if err != nil {
		t.Fatal(err.Error())
	}

	if file.Name != "Test 3" {
		t.Fatal("Le mauvais noeud à été trouvé")
	}
}

func TestSearchFileInEmptyLevel(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	_, err := searchFile("Test 3", ft.FileTree.Files)
	if err == nil {
		t.Fatal("Une erreur aurait dû survenir")
	}

	if err.Error() != "aucun fichier à ce niveau" {
		t.Fatal("La mauvaise erreur a été retournée")
	}
}

func TestSearchForAFileThatDoesNotExist(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	InsertNode(false, &ft.FileTree, "Test 1")
	InsertNode(false, &ft.FileTree, "Test 2")
	InsertNode(false, &ft.FileTree, "Test 3")
	InsertNode(false, &ft.FileTree, "Test 4")

	_, err := searchFile("Test 5", ft.FileTree.Files)
	if err == nil {
		t.Fatal("Une erreur aurait dû survenir")
	}

	if err.Error() != "le fichier n'a pas été trouvé" {
		t.Fatal("La mauvaise erreur a été retournée")
	}
}

func TestSearchForAFileInALevelWithASingleFile(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	InsertNode(false, &ft.FileTree, "Test 1")

	file, err := searchFile("Test 1", ft.FileTree.Files)
	if err != nil {
		t.Fatal(err.Error())
	}

	if file.Name != "Test 1" {
		t.Fatal("Le mauvais noeud à été trouvé")
	}
}

func TestSearchForAFileThatDoesNotInALevelWithASingleFile(t *testing.T) {
	ft := NewFileTree(&config.AppConfig{
		ConfigFile: config.ConfigFile{
			LabPath: "D:\\Projets\\test\\Lab",
		},
	})

	InsertNode(false, &ft.FileTree, "Test 1")

	_, err := searchFile("Test 5", ft.FileTree.Files)
	if err == nil {
		t.Fatal("Une erreur aurait dû survenir")
	}

	if err.Error() != "le fichier n'a pas été trouvé" {
		t.Fatal("La mauvaise erreur a été retournée")
	}
}
