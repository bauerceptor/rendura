// Example showing how to draw shapes and use a mouse.
package main

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_mouse"
	"math"
)

const (
	shapeColor = 15
	textColor  = 2
)

var (
	//go:embed sprite-sheet.png
	spriteSheetPNG []byte

	drawShape func(start, stop rendura.Position)

	drawFunctions = []func(start, stop rendura.Position){
		func(start, stop rendura.Position) {
			rendura.Rect(start.X, start.Y, stop.X, stop.Y)
			command := fmt.Sprintf("rendura.Rect(%d,%d,%d,%d)", start.X, start.Y, stop.X, stop.Y)
			printCmd(command)
		},
		func(start, stop rendura.Position) {
			rendura.RectFill(start.X, start.Y, stop.X, stop.Y)
			command := fmt.Sprintf("rendura.RectFill(%d,%d,%d,%d)", start.X, start.Y, stop.X, stop.Y)
			printCmd(command)
		},
		func(start, stop rendura.Position) {
			rendura.Line(start.X, start.Y, stop.X, stop.Y)
			command := fmt.Sprintf("rendura.Line(%d,%d,%d,%d)", start.X, start.Y, stop.X, stop.Y)
			printCmd(command)
		},
		func(start, stop rendura.Position) {
			r := radius(start.X, start.Y, stop.X, stop.Y)
			rendura.Circ(start.X, start.Y, r)

			command := fmt.Sprintf("rendura.Circ(%d,%d,%d)", start.X, start.Y, r)
			printCmd(command)
		},
		func(start, stop rendura.Position) {
			r := radius(start.X, start.Y, stop.X, stop.Y)
			rendura.CircFill(start.X, start.Y, r)
			command := fmt.Sprintf("rendura.CircFill(%d,%d,%d)", start.X, start.Y, r)
			printCmd(command)
		},
	}

	currentShapeIdx = 0

	shapeStart rendura.Position

	cursorSprites []rendura.Sprite
)

func main() {
	rendura.SetScreenSize(128, 128)
	rendura.SetTPS(60) // rendura.Update and rendura.Draw wil be executed 60 times per second

	rendura.Palette = rendura.DecodePalette(spriteSheetPNG)
	spriteSheet := rendura.DecodeCanvas(spriteSheetPNG)

	// create cursors sprite array for each shape
	for i := range drawFunctions {
		cursorSprite := rendura.SpriteFrom(spriteSheet, i*8, 0, 8, 8)
		cursorSprites = append(cursorSprites, cursorSprite)
	}

	rendura.Draw = func() {
		rendura.Cls()

		// change the shape if right mouse button was just pressed
		if rendura_mouse.Duration(rendura_mouse.Right) == 1 {
			currentShapeIdx++
			if currentShapeIdx == len(drawFunctions) {
				currentShapeIdx = 0
			}
		}

		// set initial coordinates on start dragging
		if rendura_mouse.Duration(rendura_mouse.Left) > 0 && drawShape == nil {
			shapeStart = rendura_mouse.Position
			drawShape = drawFunctions[currentShapeIdx]

		}

		// set coordinates during dragging
		if drawShape != nil {
			stop := rendura_mouse.Position
			rendura.SetColor(shapeColor)
			drawShape(shapeStart, stop)
		}

		if rendura_mouse.Duration(rendura_mouse.Left) == 0 {
			drawShape = nil
		}

		drawMousePointer()
	}

	ebiten.SetCursorMode(ebiten.CursorModeHidden) // hide cursor in Ebitengine
	rendura_ebiten.Run()
}

func drawMousePointer() {
	rendura.DrawSprite(cursorSprites[currentShapeIdx], rendura_mouse.Position.X, rendura_mouse.Position.Y)
}

func radius(x0, y0, x1, y1 int) int {
	dx := math.Abs(float64(x0 - x1))
	dy := math.Abs(float64(y0 - y1))
	return int(math.Max(dx, dy))
}

func printCmd(command string) {
	rendura.SetColor(textColor)
	rendura_cofont.Print(command, 8, 128-6-6)
}
