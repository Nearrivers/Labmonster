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
	return &AppConfig{
		Logger: logger.NewDefaultLogger(),
	}
}

func (ac *AppConfig) SetContext(ctx context.Context) {
	ac.Ctx = ctx
}

func (ac *AppConfig) GetConfigFile() ConfigFile {
	return ac.ConfigFile
}

func (ac *AppConfig) SetConfigFile(cfg ConfigFile) {
	ac.ConfigFile = cfg
}

func (ac *AppConfig) OpenCreateLabDialog() string {
	pwd, err := os.Getwd()
	if err != nil {
		ac.Logger.Error(err.Error())
		return ""
	}

	dir, err := runtime.OpenDirectoryDialog(ac.Ctx, runtime.OpenDialogOptions{
		DefaultDirectory: pwd,
		Title:            "Emplacement",
	})
	if err != nil {
		ac.Logger.Error(err.Error())
		return ""
	}

	return dir
}

func (ac *AppConfig) CreateAppConfig(configDirPath string) {
	config := ConfigFile{
		LabPath: configDirPath,
	}

	go func() {
		data, err := toml.Marshal(config)
		if err != nil {
			ac.Logger.Error(err.Error())
			return
		}

		err = os.WriteFile("config.toml", data, 0o644)
		if err != nil {
			ac.Logger.Error(err.Error())
			return
		}

		// Création du dossier "Lab" s'il n'existe pas
		if _, err := os.Stat(filepath.Join(configDirPath, "Lab")); errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(filepath.Join(configDirPath, "Lab"), 0o644)
			if err != nil {
				ac.Logger.Error(err.Error())
			}
		}
	}()

	ac.SetConfigFile(config)
}

func (ac *AppConfig) CheckConfigPresence() bool {
	if _, err := os.Stat("config.toml"); errors.Is(err, os.ErrNotExist) {
		ac.Logger.Error(err.Error())
		return false
	}
	go ac.LoadConfigFileInMemory()
	return true
}

func (ac *AppConfig) LoadConfigFileInMemory() {
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