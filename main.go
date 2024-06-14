package main

import (
	"context"
	"embed"

	"flow-poc/backend/config"
	"flow-poc/backend/topmenu"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	topmenu := topmenu.NewTopMenu()
	config := config.NewAppConfig()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "LabMonster",
		Width:            1024,
		Height:           768,
		AlwaysOnTop:      false,
		WindowStartState: options.Maximised,
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
		},
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
