package config

import (
	"context"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConfigFile struct {
	LabPath string `toml:"labpath"`
}

type AppConfig struct {
	ctx context.Context
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (ac *AppConfig) SetContext(ctx context.Context) {
	ac.ctx = ctx
}

func (ac *AppConfig) OpenCreateLabDialog() string {
	pwd, err := os.Getwd()
	if err != nil {
		logger.NewDefaultLogger().Error(err.Error())
		return ""
	}

	dir, err := runtime.OpenDirectoryDialog(ac.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: pwd,
		Title:            "Emplacement",
	})
	if err != nil {
		logger.NewDefaultLogger().Error(err.Error())
		return ""
	}

	return dir
}

func (ac *AppConfig) CreateAppConfig(configDirPath string) {
	config := ConfigFile{
		LabPath: configDirPath,
	}

	data, err := toml.Marshal(config)
	if err != nil {
		logger.NewDefaultLogger().Error(err.Error())
		return
	}

	err = os.WriteFile("config.toml", data, 0o644)
	if err != nil {
		logger.NewDefaultLogger().Error(err.Error())
		return
	}

	os.Mkdir(fmt.Sprintf("%s/Lab", configDirPath), 0o644)
}
