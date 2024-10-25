package db

import (
	"database/sql"
	repository "flow-poc/backend/db/repository"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func createDbFile() {
	_, err := os.Stat(filepath.Join("backend", "db", "core.db"))
	if err == nil || !os.IsNotExist(err) {
		return
	}

	f, errCreate := os.Create(filepath.Join("backend", "db", "core.db"))
	if errCreate != nil {
		log.Fatalf("Impossible de créer le fichier de base de donnée: %v\n", errCreate)
	}

	f.Close()
}

func ConnectToDb() *repository.Queries {
	createDbFile()

	db, err := sql.Open("sqlite3", filepath.Join("backend", "db", "core.db"))
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de donnée: %v\n", err)
	}

	return repository.New(db)
}