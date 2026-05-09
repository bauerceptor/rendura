// Example of using the high-level audio API in rendura.
//
// This program plays a sample at different pitches depending on the key pressed.
package main

import (
	_ "embed"
	"math"
	"slices"

	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_audio"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_key"
)

var (
	//go:embed "piano.raw"
	sampleRAW []byte
	//go:embed "piano.png"
	pianoPNG []byte
)

func main() {
	// Decode a raw sample file (8-bit mono, no header, no compression).
	sample := rendura_audio.DecodeRaw(sampleRAW, 16726)

	var selectedKey rendura_key.Key

	// Load palette and canvas from PNG file
	rendura.Palette = rendura.DecodePalette(pianoPNG)
	pianoCanvas := rendura.DecodeCanvas(pianoPNG)

	rendura.SetScreenSize(113, 46)
	rendura.SetTransparency(0, false) // disable transparency for color 0

	rendura.Init = func() {
		// The sample must be loaded before use,
		// but communication with the audio backend is only possible after starting the game.
		rendura_audio.LoadSample(sample)
	}

	rendura.Update = func() {
		// Check if any of the piano keys were just pressed
		for key, pitch := range buttonPitch {
			if rendura_key.Duration(key) == 1 {
				selectedKey = key
				// Play the sample on two channels: left and right — for stereo effect
				ch := rendura_audio.Chan1 | rendura_audio.Chan2
				vol := 1.0
				rendura_audio.Play(ch, sample, pitch, vol)
				break
			}
		}
	}

	rendura.Draw = func() {
		rendura.Cls()

		// Map each key color to either white (tone) or black (semitone)
		for key, color := range keyColors {
			if slices.Contains(semitoneLetters, key) {
				rendura.RemapColor(color, 0)
			} else {
				rendura.RemapColor(color, 7)
			}
		}

		// Highlight the last pressed key in gray
		if selectedKey != "" {
			rendura.RemapColor(keyColors[selectedKey], 6)
		}

		// Draw the piano image with updated color tables
		rendura.DrawCanvas(pianoCanvas, 0, 0)

		// Draw labels for each piano key
		printLetters()
	}

	rendura_ebiten.Run()
}

// cPitch = 1.0 is the base pitch (e.g. middle C).
// Change to 2.0 to play one octave higher, or 0.5 for one octave lower.
const cPitch = 1.0

// Maps keyboard keys to pitch multipliers.
var buttonPitch = map[rendura_key.Key]float64{
	rendura_key.Z:     cPitch,          // C
	rendura_key.S:     adjustPitch(1),  // C#
	rendura_key.X:     adjustPitch(2),  // D
	rendura_key.D:     adjustPitch(3),  // D#
	rendura_key.C:     adjustPitch(4),  // E
	rendura_key.V:     adjustPitch(5),  // F
	rendura_key.G:     adjustPitch(6),  // F#
	rendura_key.B:     adjustPitch(7),  // G
	rendura_key.H:     adjustPitch(8),  // G#
	rendura_key.N:     adjustPitch(9),  // A
	rendura_key.J:     adjustPitch(10), // A#
	rendura_key.M:     adjustPitch(11), // H
	rendura_key.Comma: adjustPitch(12), // C
}

func adjustPitch(i int) float64 {
	return cPitch * math.Pow(2, float64(i)/12.0)
}

// Layout of tone and semitone key labels on the piano image.
var (
	toneLetters     = []rendura_key.Key{"Z", "X", "C", "V", "B", "N", "M", ","}
	semitoneLetters = []rendura_key.Key{"S", "D", "", "G", "H", "J"}
)

const keyWidth = 13

func printLetters() {
	rendura.SetColor(16)
	for i, letter := range toneLetters {
		rendura_cofont.Print(string(letter), 9+i*keyWidth, 31)
	}

	for i, letter := range semitoneLetters {
		rendura_cofont.Print(string(letter), 16+i*keyWidth, 13)
	}
}

// Each key on the image uses a unique color
var keyColors = map[rendura_key.Key]rendura.Color{
	rendura_key.Z:     1,
	rendura_key.S:     2,
	rendura_key.X:     3,
	rendura_key.D:     4,
	rendura_key.C:     6,
	rendura_key.V:     8,
	rendura_key.G:     9,
	rendura_key.B:     10,
	rendura_key.H:     11,
	rendura_key.N:     12,
	rendura_key.J:     13,
	rendura_key.M:     14,
	rendura_key.Comma: 15,
}
