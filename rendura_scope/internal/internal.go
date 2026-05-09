package internal

import (
	"fmt"
	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_debug"
	"github.com/bauerceptor/rendura/rendura_event"
	"github.com/bauerceptor/rendura/rendura_gui"
	"github.com/bauerceptor/rendura/rendura_loop"
	"github.com/bauerceptor/rendura/rendura_mouse"
)

var bgColor, fgColor *rendura.Color

var consoleMode, pauseOnNextFrame bool

// Start launches the developer tools.
//
// Obecnie piscope wymaga, żeby gra miała rozdzielczość conajmniej 128
// pikseli w poziomie oraz 16 pikseli w pionie. Dodatkowa paleta gry musi
// używać conajmniej 2 kolorów.
//
// Pressing Ctrl+Shift+I will activate the tools in the game
func Start(backgroundColor, foregroundColor *rendura.Color) {
	bgColor = backgroundColor
	fgColor = foregroundColor

	// TODO Handle screen size change event and redraw entire gui.

	smallFont := rendura_cofont.Sheet

	registerShortcuts()

	gui := rendura_gui.New()
	attachToolbar(gui)

	rendura_loop.DebugTarget().Subscribe(rendura_loop.EventUpdate, func(rendura_loop.Event, rendura_event.Handler) {
		if consoleMode {
			gui.Update()

			if !rendura_debug.Paused() {
				theScreenRecorder.Save()
			}

			if pauseOnNextFrame {
				rendura_debug.SetPaused(true)
				pauseOnNextFrame = false
			}

			handleInputInConsoleMode()
		}
	})

	rendura_loop.DebugTarget().Subscribe(rendura_loop.EventLateDraw, func(rendura_loop.Event, rendura_event.Handler) {
		if consoleMode {
			gui.Draw()

			screen := rendura.Screen()

			y := screen.H() - smallFont.Height - 1

			prev := rendura.SetColor(*bgColor)
			defer rendura.SetColor(prev)

			pixelColor := rendura.GetPixel(rendura_mouse.Position.X, rendura_mouse.Position.Y)
			if pixelColor != *bgColor {
				rendura.SetColor(pixelColor)
			} else {
				rendura.SetColor(1)
			}
			msg := fmt.Sprintf("%d(%d,%d)", pixelColor, rendura_mouse.Position.X, rendura_mouse.Position.Y)
			smallFont.Print(msg, 50, y+2)
		}
	})
}
