package internal

import (
	_ "embed"
	"github.com/bauerceptor/rendura"
)

//go:embed "icons.png"
var iconsPNG []byte

var icons = struct {
	AlignTop    rendura.Sprite
	AlignBottom rendura.Sprite
	Screen      rendura.Sprite
	Palette     rendura.Sprite
	ColorTables rendura.Sprite
	Variables   rendura.Sprite
	Paint       rendura.Sprite
	Separator   rendura.Sprite
	Snap        rendura.Sprite
	Prev        rendura.Sprite
	Pause       rendura.Sprite
	Play        rendura.Sprite
	Next        rendura.Sprite
	Exit        rendura.Sprite
}{}

func init() {
	prevPalette := rendura.Palette
	defer func() {
		rendura.Palette = prevPalette
	}()

	rendura.Palette = rendura.DecodePalette(iconsPNG)
	iconsSheet := rendura.DecodeCanvas(iconsPNG)

	icons.AlignTop = rendura.SpriteFrom(iconsSheet, 0, 0, 8, 8)
	icons.AlignBottom = rendura.SpriteFrom(iconsSheet, 8, 0, 8, 8)
	icons.Screen = rendura.SpriteFrom(iconsSheet, 16, 0, 8, 8)
	icons.Palette = rendura.SpriteFrom(iconsSheet, 24, 0, 8, 8)
	icons.ColorTables = rendura.SpriteFrom(iconsSheet, 32, 0, 8, 8)
	icons.Variables = rendura.SpriteFrom(iconsSheet, 40, 0, 8, 8)
	icons.Paint = rendura.SpriteFrom(iconsSheet, 48, 0, 8, 8)
	icons.Separator = rendura.SpriteFrom(iconsSheet, 58, 0, 4, 8)
	icons.Snap = rendura.SpriteFrom(iconsSheet, 64, 0, 8, 8)
	icons.Prev = rendura.SpriteFrom(iconsSheet, 74, 0, 5, 8)
	icons.Pause = rendura.SpriteFrom(iconsSheet, 82, 0, 5, 8)
	icons.Play = rendura.SpriteFrom(iconsSheet, 90, 0, 5, 8)
	icons.Next = rendura.SpriteFrom(iconsSheet, 98, 0, 5, 8)
	icons.Exit = rendura.SpriteFrom(iconsSheet, 112, 0, 8, 8)
}
