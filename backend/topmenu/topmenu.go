package topmenu

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TopMenu struct {
	isMaximized bool
	ctx context.Context
}

func NewTopMenu() *TopMenu {
	return &TopMenu{
		isMaximized: true,
	}
}

func (tm *TopMenu) SetContext(ctx context.Context) {
	tm.ctx = ctx
}

func (tm *TopMenu) QuitApp() {
	runtime.Quit(tm.ctx)
}

func (tm *TopMenu) MaximiseOrUnmaximiseWindow() {
	runtime.WindowToggleMaximise(tm.ctx)
}

func (tm *TopMenu) MinimiseWindow() {
		runtime.WindowMinimise(tm.ctx)
}
