package config

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConfigFile struct {
	LabPath string `toml:"labpath"`
}

type AppConfig struct {
	Ctx        context.Context
	Logger     logger.Logger
	ConfigFile ConfigFile
}

func NewAppConfig() *AppConfig {
	ac := AppConfig{
		Logger: logger.NewDefaultLogger(),
	}
	ac.CheckConfigPresenceAndLoadIt()

	return &ac
}

func (ac *AppConfig) SetContext(ctx context.Context) {
	ac.Ctx = ctx
}

func (ac *AppConfig) SetConfigFile(cfg ConfigFile) {
	ac.ConfigFile = cfg
}

func (ac *AppConfig) OpenCreateLabDialog() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		ac.Logger.Error(err.Error())
		return "", err
	}

	dir, err := runtime.OpenDirectoryDialog(ac.Ctx, runtime.OpenDialogOptions{
		DefaultDirectory: pwd,
		Title:            "Emplacement",
	})
	if err != nil {
		ac.Logger.Error(err.Error())
		return "", err
	}

	return dir, nil
}

func (ac *AppConfig) CreateAppConfig(configDirPath string) {
	config := ConfigFile{
		LabPath: configDirPath,
	}

	go func() {
		// Création du dossier "Lab" s'il n'existe pas
		err := os.MkdirAll(filepath.Join(configDirPath, ".labmonster"), os.ModePerm)
		if err != nil {
			ac.Logger.Error(err.Error())
		}
	}()

	data, err := toml.Marshal(config)
	if err != nil {
		ac.Logger.Error(err.Error())
		return
	}

	err = os.WriteFile("config.toml", data, os.ModePerm)
	if err != nil {
		ac.Logger.Error(err.Error())
		return
	}

	ac.SetConfigFile(config)
}

// Vérifie la présence du fichie de configuration et le charge si c'est le cas
func (ac *AppConfig) CheckConfigPresenceAndLoadIt() bool {
	if _, err := os.Stat("config.toml"); errors.Is(err, os.ErrNotExist) {
		ac.Logger.Error(err.Error())
		return false
	}
	go ac.LoadConfigFile()
	return true
}

// Charge le fichier de configuration
func (ac *AppConfig) LoadConfigFile() {
	f, err := os.Open("config.toml")
	if err != nil {
		ac.Logger.Error(err.Error())
		return
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		ac.Logger.Error(err.Error())
		return
	}

	cfg := ConfigFile{}
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		ac.Logger.Error(err.Error())
		return
	}

	ac.SetConfigFile(cfg)
}
