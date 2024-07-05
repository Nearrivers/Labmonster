package filetree

import (
	"errors"
	"flow-poc/backend/config"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateNewFile(t *testing.T) {
	t.Run("Création fichier basique", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		defer ft.DeleteFile("Sans titre.json")

		_, err := ft.CreateNewFile()
		if err != nil {
			t.Errorf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
		}

		_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json"))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Error("Le fichier n'a pas été créé")
			}
			t.Errorf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
		}

		if errors.Is(err, os.ErrNotExist) {
			t.Error("Le fichier n'a pas été créé")
		}
	})

	t.Run("Création de plusieurs fichiers à la suite", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		defer ft.DeleteFile("Sans titre.json")
		defer ft.DeleteFile("Sans titre 1.json")
		defer ft.DeleteFile("Sans titre 2.json")

		cpt := 0
		for cpt < 3 {
			_, err := ft.CreateNewFile()
			if err != nil {
				t.Errorf("Une erreur est survenue lors de la création du fichier: %v", err.Error())
			}

			if cpt > 0 {
				_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, fmt.Sprintf("Sans titre %d.json", cpt)))
			} else {
				_, err = os.Stat(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json"))
			}

			if errors.Is(err, os.ErrNotExist) {
				t.Error("Le fichier n'a pas été créé")
			} else {
				t.Errorf("Une erreur est survenue lors de la création du fichier: %q", err)
			}

			cpt++
		}
	})
}

// Test du happy path
func TestSearchFile(t *testing.T) {
	t.Run("Recherche d'un noeud existant", func(t *testing.T) {
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
			t.Error(err.Error())
		}

		if file.Name != "Test 3" {
			t.Error("Le mauvais noeud à été trouvé")
		}
	})

	t.Run("Recherche d'un noeud dans un niveau vide", func(t *testing.T) {
		ft := NewFileTree(&config.AppConfig{
			ConfigFile: config.ConfigFile{
				LabPath: "D:\\Projets\\test\\Lab",
			},
		})

		_, err := searchFile("Test 3", ft.FileTree.Files)
		if err == nil {
			t.Error("Une erreur aurait dû survenir")
		}

		if err != ErrNoFileInThisLevel {
			t.Error("La mauvaise erreur a été retournée")
		}
	})

	t.Run("Recherche d'un noeud qui n'existe pas", func(t *testing.T) {
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
			t.Error("Une erreur aurait dû survenir")
		}

		if err != ErrNodeNotFound {
			t.Error("La mauvaise erreur a été retournée")
		}
	})
}
