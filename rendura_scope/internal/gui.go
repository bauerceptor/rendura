package internal

import (
	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_debug"
	"github.com/bauerceptor/rendura/rendura_gui"
)

func attachToolbar(parent *rendura_gui.Element) *rendura_gui.Element {
	toolbar := rendura_gui.Attach(parent, 0, rendura.Screen().H()-9, rendura.Screen().W(), 9)
	toolbar.OnDraw = func(event rendura_gui.DrawEvent) {
		prev := rendura.SetColor(*bgColor)
		defer rendura.SetColor(prev)
		rendura.RectFill(0, 0, toolbar.W, toolbar.H)
	}

	// attachIconButton(toolbar, icons.AlignTop, 0) // icon hidden until implemented
	// attachIconButton(toolbar, icons.Screen, 8) // icon hidden for now because screen inspector is the only tab
	// attachIconButton(toolbar, icons.Palette, 16) // icon hidden until implemented
	// attachIconButton(toolbar, icons.Variables, 24) // icon hidden until implemented
	// attachIconButton(toolbar, icons.Paint, 32) // icon hidden until implemented

	snap := attachIconButton(toolbar, icons.Snap, rendura.Screen().W()-34)
	snap.OnTap = func(event rendura_gui.Event) {
		captureSnapshot()
	}

	prev := attachIconButton(toolbar, icons.Prev, rendura.Screen().W()-24)
	prev.OnTap = func(event rendura_gui.Event) {
		showPrevSnapshot()
	}
	prev.OnUpdate = func(rendura_gui.UpdateEvent) {
		if theScreenRecorder.HasPrev() {
			prev.Icon = icons.Prev
		} else {
			prev.Icon = rendura.Sprite{}
		}
	}

	playPause := attachIconButton(toolbar, icons.Pause, rendura.Screen().W()-19)
	playPause.OnTap = func(rendura_gui.Event) {
		pauseOrResume()
	}
	playPause.OnUpdate = func(rendura_gui.UpdateEvent) {
		if rendura_debug.Paused() {
			playPause.Icon = icons.Pause
		} else {
			playPause.Icon = icons.Play
		}
	}

	next := attachIconButton(toolbar, icons.Next, rendura.Screen().W()-14)
	next.OnTap = func(event rendura_gui.Event) {
		showNextSnapshot()
	}

	exit := attachIconButton(toolbar, icons.Exit, rendura.Screen().W()-8)
	exit.OnTap = func(event rendura_gui.Event) {
		exitConsoleMode()
	}

	return toolbar
}

type IconButton struct {
	*rendura_gui.Element

	Icon rendura.Sprite
}

func attachIconButton(parent *rendura_gui.Element, icon rendura.Sprite, x int) *IconButton {
	btn := rendura_gui.Attach(parent, x, 0, icon.W, icon.H+1)
	iconBtn := &IconButton{Icon: icon, Element: btn}
	btn.OnDraw = func(event rendura_gui.DrawEvent) {
		y := 0
		if event.Pressed {
			y = 1
		}

		prevColorTable := rendura.ColorTables[0]
		defer func() {
			rendura.ColorTables[0] = prevColorTable
		}()
		rendura.RemapColor(0, *bgColor) // 0 is bg color in icons.png
		rendura.RemapColor(1, *fgColor) // 1 is fg color in icons.png
		rendura.DrawSprite(iconBtn.Icon, 0, y)
	}
	return iconBtn
}
