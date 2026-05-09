package main

import (
	_ "embed"
	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_pad"
)

//go:embed "gamepad.png"
var gamepadPNG []byte

func main() {
	rendura.SetScreenSize(85, 60)
	rendura.Palette = rendura.DecodePalette(gamepadPNG)
	gamepad := rendura.DecodeCanvas(gamepadPNG)

	buttonSprites := map[rendura_pad.Button]rendura.Sprite{
		rendura_pad.X:      rendura.SpriteFrom(gamepad, 48, 18, 9, 9),
		rendura_pad.Y:      rendura.SpriteFrom(gamepad, 58, 10, 9, 9),
		rendura_pad.B:      rendura.SpriteFrom(gamepad, 68, 18, 9, 9),
		rendura_pad.A:      rendura.SpriteFrom(gamepad, 58, 26, 9, 9),
		rendura_pad.Left:   rendura.SpriteFrom(gamepad, 11, 19, 8, 8),
		rendura_pad.Right:  rendura.SpriteFrom(gamepad, 25, 19, 8, 8),
		rendura_pad.Top:    rendura.SpriteFrom(gamepad, 19, 13, 6, 8),
		rendura_pad.Bottom: rendura.SpriteFrom(gamepad, 19, 25, 6, 8),
	}

	rendura.Draw = func() {
		rendura.Cls()
		rendura.DrawCanvas(gamepad, 0, 0)

		for button, sprite := range buttonSprites {
			if rendura_pad.Duration(button) > 0 { // duration is > 0 when button is pressed
				rendura.DrawSprite(sprite, sprite.X, sprite.Y+1) // draw pressed button
			}
		}

		rendura_cofont.Print("PRESS BTN ON GAMEPAD", 3, 50)
	}

	rendura_ebiten.Run()
}
