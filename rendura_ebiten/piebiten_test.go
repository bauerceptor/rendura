//go:build !js

package rendura_ebiten_test

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_ebiten/internal/ebitentesting"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	ebitentesting.MainWithRunLoop(m)
}

func TestCopyCanvasToEbitenImage(t *testing.T) {
	t.Run("should copy canvas to Ebiten image", func(t *testing.T) {
		canvas := rendura.NewCanvas(3, 2)
		rendura.Palette[0] = 0xFF0000
		rendura.Palette[1] = 0x00FF00
		rendura.Palette[2] = 0x0000FF
		rendura.Palette[3] = 0x00FFFF
		rendura.Palette[4] = 0xFFFF00
		rendura.Palette[5] = 0xFF00FF
		canvas.SetAll(
			0, 1, 2,
			3, 4, 5,
		)
		img := ebiten.NewImage(3, 2)
		// when
		rendura_ebiten.CopyCanvasToEbitenImage(canvas, img)
		// then
		bounds := img.Bounds()
		out := make([]byte, bounds.Dx()*bounds.Dy()*4)
		img.ReadPixels(out)
		for i := 0; i < 6; i++ {
			pix := out[i*4 : i*4+4]
			rgb := rendura.FromRGB(pix[0], pix[1], pix[2])
			assert.Equal(t, rendura.Palette[i], rgb, "pixel %d", i)
			assert.Equal(t, uint8(0xff), pix[3], "alpha for pixel %d", i)
		}
	})

	t.Run("should take into account palette mapping", func(t *testing.T) {
		canvas := rendura.NewCanvas(1, 1)
		rendura.Palette[0] = 0x000000
		rendura.Palette[1] = 0xFFFFFF
		rendura.PaletteMapping[0] = 1 // map 0 to 1 (0x000000 to 0xFFFFFF)
		img := ebiten.NewImage(1, 1)
		// when
		rendura_ebiten.CopyCanvasToEbitenImage(canvas, img)
		// then
		bounds := img.Bounds()
		out := make([]byte, bounds.Dx()*bounds.Dy()*4)
		img.ReadPixels(out)
		pix := out[0:4]
		rgb := rendura.FromRGB(pix[0], pix[1], pix[2])
		assert.Equal(t, rendura.RGB(0xFFFFFF), rgb)
	})
}
