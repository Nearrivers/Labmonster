package main

import (
	"context"
	"embed"
	"log"
	"time"

	"flow-poc/backend/config"
	dirhandler "flow-poc/backend/filesystem/dir_handler"
	"flow-poc/backend/filesystem/file_handler"
	"flow-poc/backend/filesystem/node"
	"flow-poc/backend/topmenu"
	"flow-poc/backend/watcher"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type Bar struct {
	Name string
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	topmenu := topmenu.NewTopMenu()
	config := config.NewAppConfig()
	fh := file_handler.NewFileHandler(config)
	dh := dirhandler.NewDirHandler(config, fh.RecentFiles)
	w := watcher.New(config)

	go func() {
		w.Wait()
	}()

	go func() {
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		for {
			select {
			case err := <-w.Error:
				log.Fatalln(err)
			case evt := <-w.Event:
				evt.MarshalFrontend(config.ConfigFile.LabPath)
				log.Printf("event reçu %s", evt)
				runtime.EventsEmit(w.Ctx, "fsop", evt)
			}
		}
	}()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "LabMonster",
		Width:            1024,
		Height:           768,
		AlwaysOnTop:      false,
		WindowStartState: options.Maximised,
		Frameless:        true,
		DisableResize:    false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			topmenu.SetContext(ctx)
			config.SetContext(ctx)
			w.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			topmenu,
			config,
			fh,
			dh,
		},
		EnumBind: []interface{}{
			watcher.FsOps,
			node.FTypes,
			node.DTypes,
		},
		OnShutdown: func(ctx context.Context) {
			fh.RecentFiles.SaveRecentlyOpended()
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
