// Package pifont provides functionality for rendering text
// using bitmap fonts.
package rendura_font

import (
	"github.com/bauerceptor/rendura/pkg/rendura"
)

// Sheet is a character sheet used for rendering text.
type Sheet struct {
	Chars   map[rune]rendura.Sprite
	Height  int
	FgColor rendura.Color // font color on sprites
	BgColor rendura.Color // background color on sprites
}

var intermediateCanvas rendura.Canvas // text is first rendered here to change its color from FgColor to selected color

var prevFgColorTable [rendura.MaxColors]rendura.Color
var prevBgColorTable [rendura.MaxColors]rendura.Color

// Print draws text using the current draw color.
//
// Returns the x, y position where you can continue writing text.
func (s Sheet) Print(str string, x, y int) (currentX, currentY int) {
	originalDrawTarget := rendura.DrawTarget()
	if intermediateCanvas.W() != originalDrawTarget.W() || intermediateCanvas.H() != originalDrawTarget.H() {
		intermediateCanvas = rendura.NewCanvas(originalDrawTarget.W(), originalDrawTarget.H())
	}

	currentColor := rendura.GetColor()

	prevFgColorTable = rendura.ColorTables[0][s.FgColor]
	prevBgColorTable = rendura.ColorTables[0][s.BgColor]

	// create fake bg color to avoid a situation when fg and bg colors are the same
	bgColor := (currentColor + 1) % rendura.MaxColors
	intermediateCanvas.Clear(s.BgColor)
	rendura.RemapColor(s.FgColor, currentColor)
	rendura.RemapColor(s.BgColor, bgColor)
	rendura.SetDrawTarget(intermediateCanvas)

	// first draw text in selected color on intermediateCanvas
	currentX, currentY = s.PrintOriginal(str, x, y)

	// revert color tables
	rendura.ColorTables[0][s.FgColor] = prevFgColorTable
	rendura.ColorTables[0][s.BgColor] = prevBgColorTable

	// make bgColor transparent
	prevBgColorTable = rendura.ColorTables[0][bgColor]
	rendura.SetTransparency(bgColor, true)

	// now copy text in target color on original draw target
	coloredText := rendura.Sprite{
		Area: rendura.Area[int]{
			X: x - rendura.Camera.X,
			Y: y - rendura.Camera.Y,
			W: currentX - x,
			H: currentY - y + s.Height,
		},
		Source: intermediateCanvas,
	}
	rendura.SetDrawTarget(originalDrawTarget)
	rendura.DrawSprite(coloredText, x, y)

	// revert bgColor transparency
	rendura.ColorTables[0][bgColor] = prevBgColorTable

	return
}

// PrintOriginal prints the text using its original colors.
func (s Sheet) PrintOriginal(str string, x, y int) (maxX, currentY int) {
	maxX = x
	currentX := x
	currentY = y
	for _, r := range str {
		if r == '\n' {
			currentX = x
			currentY += s.Height
			continue
		}
		sprite := s.Chars[r]
		rendura.DrawSprite(sprite, currentX, currentY)
		currentX += sprite.W
		maxX = max(maxX, currentX)
	}

	return
}

// PrintStroked prints the text with a stroke effect.
//
// The text is drawn using the specified foreground and stroke colors.
func (s Sheet) PrintStroked(text string, x, y int, fgColor, strokeColor rendura.Color) (currentX, currentY int) {
	prevColor := rendura.SetColor(strokeColor)
	for l := y - 1; l <= y+1; l++ {
		s.Print(text, x-1, l)
		s.Print(text, x, l)
		s.Print(text, x+1, l)
	}

	rendura.SetColor(fgColor)
	currentX, currentY = s.Print(text, x, y)

	rendura.SetColor(prevColor)

	return
}

// Size returns the dimensions of the text without rendering it to the draw target.
func (s Sheet) Size(text string) (width, height int) {
	originalDrawTarget := rendura.SetDrawTarget(intermediateCanvas)
	defer rendura.SetDrawTarget(originalDrawTarget)

	return s.PrintOriginal(text, 0, 0)
}
