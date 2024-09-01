package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow
	var moveHandle *walk.PushButton

	bgColor := walk.RGB(240, 240, 245)   // 연한 청회색
	btnColor := walk.RGB(220, 220, 230)  // 연한 보라색
	moveColor := walk.RGB(200, 200, 210) // 이동 버튼

	var moving bool
	var lastX, lastY int

	MainWindow{
		AssignTo:   &mw,
		Title:      "Arrow Copier",
		Background: SolidColorBrush{Color: bgColor},
		Layout:     VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:       "↑",
						Font:       Font{Family: "Segoe UI", PointSize: 14, Bold: true},
						MaxSize:    Size{40, 40},
						MinSize:    Size{40, 40},
						Background: SolidColorBrush{Color: btnColor},
						OnClicked: func() {
							walk.Clipboard().SetText("↑")
							mw.SetTitle("Copied: ↑")
						},
					},
					HSpacer{},
					PushButton{
						Text:       "←",
						Font:       Font{Family: "Segoe UI", PointSize: 14, Bold: true},
						MaxSize:    Size{40, 40},
						MinSize:    Size{40, 40},
						Background: SolidColorBrush{Color: btnColor},
						OnClicked: func() {
							walk.Clipboard().SetText("←")
							mw.SetTitle("Copied: ←")
						},
					},
					VSpacer{
						Size: 40,
					},
					PushButton{
						Text:       "→",
						Font:       Font{Family: "Segoe UI", PointSize: 14, Bold: true},
						MaxSize:    Size{40, 40},
						MinSize:    Size{40, 40},
						Background: SolidColorBrush{Color: btnColor},
						OnClicked: func() {
							walk.Clipboard().SetText("→")
							mw.SetTitle("Copied: →")
						},
					},
					HSpacer{},
					PushButton{
						Text:       "↓",
						Font:       Font{Family: "Segoe UI", PointSize: 14, Bold: true},
						MaxSize:    Size{40, 40},
						MinSize:    Size{40, 40},
						Background: SolidColorBrush{Color: btnColor},
						OnClicked: func() {
							walk.Clipboard().SetText("↓")
							mw.SetTitle("Copied: ↓")
						},
					},
					HSpacer{},
				},
			},
			PushButton{
				AssignTo:   &moveHandle,
				Text:       "=",
				Font:       Font{Family: "Segoe UI", PointSize: 8},
				MaxSize:    Size{120, 15},
				MinSize:    Size{120, 15},
				Background: SolidColorBrush{Color: moveColor},
			},
		},
	}.Create()

	if mw == nil {
		panic("Failed to create main window")
	}

	moveHandle.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			moving = true
			lastX, lastY = x, y
		}
	})

	moveHandle.MouseMove().Attach(func(x, y int, button walk.MouseButton) {
		if moving {
			dx, dy := x-lastX, y-lastY
			newX, newY := mw.X()+dx, mw.Y()+dy
			mw.SetX(newX)
			mw.SetY(newY)
		}
	})

	moveHandle.MouseUp().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			moving = false
		}
	})

	mw.SetSize(walk.Size{Width: 120, Height: 130})
	mw.Run()
}
