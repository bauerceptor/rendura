package internal

import (
	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_debug"
	"github.com/bauerceptor/rendura/rendura_snap"
	"log"
)

func enterConsoleMode() {
	log.Println("Entering console")
	prev := rendura.SetColor(*bgColor)
	rendura.Rect(0, 0, rendura.Screen().W()-1, rendura.Screen().H()-1)
	rendura.SetColor(prev)
	consoleMode = true
}

func exitConsoleMode() {
	log.Println("Exiting console")
	theScreenRecorder.ShowPrev()
	theScreenRecorder.Reset()
	consoleMode = false
	rendura_debug.SetPaused(false)
}

func captureSnapshot() {
	f, err := rendura_snap.CaptureOrErr()
	if err != nil {
		log.Println("Error capturing screenshot:", err)
	} else {
		log.Println("Screenshot saved to", f)
	}
}

func showPrevSnapshot() {
	rendura_debug.SetPaused(true)
	theScreenRecorder.ShowPrev()
}

func showNextSnapshot() {
	if !theScreenRecorder.ShowNext() {
		rendura_debug.SetPaused(false)
		pauseOnNextFrame = true
	}
}

func pauseOrResume() {
	if consoleMode {
		theScreenRecorder.GoToLast()

		rendura_debug.SetPaused(!rendura_debug.Paused())
	}
}
