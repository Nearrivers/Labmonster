package topmenu

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TopMenu struct {
	ctx context.Context
}

func NewTopMenu() *TopMenu {
	return &TopMenu{}
}

func (tm *TopMenu) SetContext(ctx context.Context) {
	tm.ctx = ctx
}

func (tm *TopMenu) QuitApp() {
	runtime.Quit(tm.ctx)
}
