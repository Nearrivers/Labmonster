package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func TestCheckConfigPresence(t *testing.T) {
	ac := NewAppConfig()
	wd, err := os.Getwd()
	if err != nil {
		errors.Cause(err)
		t.Fatal("Impossible de récupérer le répertoire de travail")
	}

	defer os.Remove(filepath.Join(wd, "config.toml"))
	defer os.Remove(filepath.Join(wd, ".labmonster"))

	ac.CreateAppConfig(wd)
	want := true
	result := ac.CheckConfigPresenceAndLoadIt()

	if result != want {
		t.Fatal("La config n'a pas été trouvée alors qu'elle est créée")
	}
}
