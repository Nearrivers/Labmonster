package main

import (
	"context"
	"embed"

	"flow-poc/backend/config"
	"flow-poc/backend/filetree"
	"flow-poc/backend/lab/recent"
	"flow-poc/backend/topmenu"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	maxRecentlyOpenedFiles = 15
)

func main() {
	// Create an instance of the app structure
	app := NewApp()
	topmenu := topmenu.NewTopMenu()
	config := config.NewAppConfig()
	filetree := filetree.NewFileTree(config)
	recent := recent.NewRecentlyOpened(config, maxRecentlyOpenedFiles)

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "LabMonster",
		Width:            1024,
		Height:           768,
		AlwaysOnTop:      false,
		WindowStartState: options.Maximised,
		Frameless: true,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			topmenu.SetContext(ctx)
			config.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			topmenu,
			config,
			filetree,
			recent,
		},
		OnShutdown: func(ctx context.Context) {
			recent.SaveRecentlyOpended()
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
