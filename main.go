package main

import (
	"context"
	"embed"
	"flow-poc/backend/topmenu"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	topmenu := topmenu.NewTopMenu()

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "Fighting games flowchart",
		Width:       1024,
		Height:      768,
		AlwaysOnTop: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			topmenu.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			topmenu,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
