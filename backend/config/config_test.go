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
		t.Error("cannot retrieve working directory")
	}

	defer os.Remove(filepath.Join(wd, configFileName))
	defer os.Remove(filepath.Join(wd, ".labmonster"))

	ac.CreateAppConfig(wd)
	want := true
	got := ac.CheckConfigPresenceAndLoadIt()

	if got != want {
		t.Error("config was not found")
	}
}
