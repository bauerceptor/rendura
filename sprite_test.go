package rendura_test

import (
	"testing"

	"github.com/bauerceptor/rendura"
)

func TestStretch(t *testing.T) {
	// temporary test
	dst := rendura.NewCanvas(16, 16)
	rendura.SetDrawTarget(dst)

	src := rendura.NewCanvas(8, 8)
	src.Clear(7)

	spr := rendura.CanvasSprite(src)

	rendura.Stretch(spr, 0, 0, 8, 8)
	rendura.Stretch(spr, -1, 0, 8, 8)
	rendura.Stretch(spr, 0, -1, 8, 8)
	rendura.Stretch(spr, 16, 0, 8, 8)
	rendura.Stretch(spr, 0, 16, 8, 8)

	rendura.Stretch(spr.WithFlipX(true), 0, 0, 8, 8)
	rendura.Stretch(spr.WithFlipY(true), 0, 0, 8, 8)

	rendura.Stretch(spr.WithSize(0, 0), 0, 0, 8, 8)
}
