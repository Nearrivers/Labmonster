package filetree

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Fonction utilitaire qui permet de déterminer si un fichier existe au chemin indiqué
func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Créer un fichier "Sans titre" à la racine et l'ajoute aux noeuds
func (ft *FileTreeExplorer) CreateNewFile() (string, error) {
	// Création d'un fichier "Sans titre.json" si ce dernier n'existe pas
	if !doesFileExist(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json")) {
		_, err := os.Create(filepath.Join(ft.Cfg.ConfigFile.LabPath, "Sans titre.json"))
		if err != nil {
			return "", err
		}

		newNode := InsertNode(false, &ft.FileTree, "Sans titre.json")
		SortNodes(ft.FileTree.Files)
		return newNode.Name, nil
	}

	// Si un fichier "Sans titre.json" existe déjà alors on va boucler avec un compteur afin de créer un
	// "Sans titre n.json" avec n un compteur incrémenté à chaque fois qu'un fichier du même nom est trouvé
	cpt := 1
	for {
		name := fmt.Sprintf("Sans titre %d.json", cpt)
		if doesFileExist(filepath.Join(ft.Cfg.ConfigFile.LabPath, name)) {
			cpt++
			continue
		}

		_, err := os.Create(filepath.Join(ft.Cfg.ConfigFile.LabPath, name))
		if err != nil {
			return "", err
		}

		newNode := InsertNode(false, &ft.FileTree, name)
		SortNodes(ft.FileTree.Files)
		return newNode.Name, nil
	}
}

// Sauvegarde le fichier JSON du graph
func SaveFile() {}

// Renomme le fichier
func RenameFile() {}

// Supprime le fichier
func (ft *FileTreeExplorer) DeleteFile(pathFromRoot string) error {
	err := os.Remove(filepath.Join(ft.Cfg.ConfigFile.LabPath, pathFromRoot))
	if err != nil {
		return err
	}

	return nil
}

// Déplace le fichier
func MoveFile() {}

// Créer un fichier xxxx_copie.json
func DuplicateFile() {}

// Implémente une recherche binaire du noeud via son nom
// Les noms sont uniques et triés dans l'ordre alphabétique
func searchFile(name string, level []*Node) (*Node, error) {
	if len(level) == 0 {
		return nil, errors.New("aucun fichier à ce niveau")
	}

	if len(level) == 1 {
		if level[0].Name == name {
			return level[0], nil
		} else {
			return nil, errors.New("le fichier n'a pas été trouvé")
		}
	}

	low := 0
	high := len(level) - 1

	for low <= high {
		median := (low + high) / 2
		medianName := level[median]

		if medianName.Name < name {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low != len(level) && level[low].Name == name {
		return level[low], nil
	}

	return nil, errors.New("le fichier n'a pas été trouvé")
}
