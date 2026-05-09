// Example of programming audio using the low-level API.
//
// This API can be used, for example, to implement packages
// capable of playing module formats (MOD, XM, etc.) and sound effects.
package main

import (
	_ "embed"
	"log"

	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_audio"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_event"
	"github.com/bauerceptor/rendura/rendura_loop"
	"github.com/bauerceptor/rendura/rendura_mouse"
)

//go:embed "wave.wav"
var clickWav []byte

func main() {
	sample := rendura_audio.DecodeWav(clickWav)

	rendura_cofont.Print("PRESS LEFT MOUSE BUTTON TO PLAY SFX", 90, 80)

	rendura.Init = func() {
		// The sample must be loaded before use,
		// but communication with the audio backend is only possible after starting the game.
		rendura_audio.LoadSample(sample)
	}

	// function that schedules playing the SFX
	scheduleSFX := func() {
		// All commands are scheduled with a minimum delay of 0 seconds.
		// However, this doesn't mean the sound plays instantly.
		// On desktop, the delay is around 20 ms; in browsers, about 60 ms.
		// The backend automatically delays commands to reduce audio glitches.
		delay := 0.0

		// Use two channels at once.
		// Chan1 is sent to a left speaker, Chan2 is sent to a right speaker.
		ch := rendura_audio.Chan1 | rendura_audio.Chan2

		// remove all planned commands from channels
		rendura_audio.ClearChan(ch, delay)

		// set the sample to play from the beginning (offset=0)
		rendura_audio.SetSample(ch, sample, 0, delay)

		// the sound is very short, so we need to loop it.
		// the loop covers the entire sample.
		rendura_audio.SetLoop(ch, 0, sample.Len(), rendura_audio.LoopForward, delay)

		for i := 1.0; i > -0.01; i -= 0.01 {
			// gradually reduce the volume to 0
			rendura_audio.SetVolume(ch, i, delay)

			// gradually reduce the pitch down to 0
			pitch := 1.0 - delay
			rendura_audio.SetPitch(ch, pitch, delay)
			delay += 0.01
		}
	}

	leftDown := rendura_mouse.EventButton{
		Type:   rendura_mouse.EventButtonDown,
		Button: rendura_mouse.Left,
	}
	rendura_mouse.ButtonTarget().Subscribe(leftDown, func(rendura_mouse.EventButton, rendura_event.Handler) {
		scheduleSFX()
	})

	rendura_loop.Target().Subscribe(rendura_loop.EventUpdate, func(rendura_loop.Event, rendura_event.Handler) {
		// rendura_audio.Time should be used by code that wants to, for example,
		// record when track playback started.
		log.Println("TIME", rendura_audio.Time)
	})

	rendura_ebiten.Run()
}
