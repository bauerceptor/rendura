package internal

import (
	rendura_event "github.com/bauerceptor/rendura/rendura_event"
	rendura_key "github.com/bauerceptor/rendura/rendura_key"
)

func handleInputInConsoleMode() {
	right := rendura_key.Duration(rendura_key.Right)
	if right > 0 {
		if right == 1 || right > 10 {
			showNextSnapshot()
		}
	} else {
		left := rendura_key.Duration(rendura_key.Left)
		if left == 1 || left > 10 {
			showPrevSnapshot()
		}
	}
}

func registerShortcuts() {
	// CTRL+SHIFT+I
	onCtrlShiftI := func() {
		if !consoleMode {
			enterConsoleMode()
		} else {
			exitConsoleMode()
		}
	}
	rendura_key.RegisterShortcut(onCtrlShiftI, rendura_key.Ctrl, rendura_key.Shift, rendura_key.I)

	// F12
	f12Down := rendura_key.Event{Type: rendura_key.EventDown, Key: rendura_key.F12}
	rendura_key.DebugTarget().Subscribe(f12Down, func(rendura_key.Event, rendura_event.Handler) {
		if consoleMode {
			captureSnapshot()
		}
	})

	// Space
	spaceDown := rendura_key.Event{Type: rendura_key.EventDown, Key: rendura_key.Space}
	rendura_key.DebugTarget().Subscribe(spaceDown, func(rendura_key.Event, rendura_event.Handler) {
		pauseOrResume()
	})

	// Esc
	escDown := rendura_key.Event{Type: rendura_key.EventDown, Key: rendura_key.Esc}
	rendura_key.DebugTarget().Subscribe(escDown, func(rendura_key.Event, rendura_event.Handler) {
		if consoleMode {
			exitConsoleMode()
		}
	})
}
