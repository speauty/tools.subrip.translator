package tt_ui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"runtime"
	"time"
)

func (customG *Gui) defaultStatusBars() []StatusBarItem {
	mousePoint := new(win.POINT)
	win.GetCursorPos(mousePoint)
	var mouseBar *walk.StatusBarItem
	var pageBar *walk.StatusBarItem
	go func() {
		ticker := time.Tick(time.Millisecond * 100)
		for true {
			select {
			case <-ticker:
				win.GetCursorPos(mousePoint)
				if mouseBar != nil {
					_ = mouseBar.SetText(fmt.Sprintf("Mouse: [%4d, %4d]", mousePoint.X, mousePoint.Y))
				}
				if customG.pageCtl.current != nil && pageBar != nil {
					_ = pageBar.SetText(fmt.Sprintf("Page: %4s", customG.pageCtl.current.GetName()))
				}
				ticker = time.Tick(time.Millisecond * 200)
			default:
				continue
			}

		}
	}()

	return []StatusBarItem{
		{Text: fmt.Sprintf("OS: %s", runtime.GOOS), Width: 70},
		{
			AssignTo: &mouseBar, Text: fmt.Sprintf("Mouse: [%3d, %3d]", mousePoint.X, mousePoint.Y), Width: 105,
		},
		{AssignTo: &pageBar, Text: "Page: 暂无页面", Width: 85},
	}
}
