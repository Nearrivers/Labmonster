package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckConfigPresence(t *testing.T) {
	ac := NewAppConfig()
	wd, err := os.Getwd()
	if err != nil {
		t.Error("Impossible de récupérer le répertoire de travail")
	}

	defer os.Remove(filepath.Join(wd, "config.toml"))
	defer os.Remove(filepath.Join(wd, ".labmonster"))

	ac.CreateAppConfig(wd)
	want := true
	got := ac.CheckConfigPresenceAndLoadIt()

	if got != want {
		t.Error("La config n'a pas été trouvée alors qu'elle est créée")
	}
}
