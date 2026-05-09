package rendura_snap_test

import (
	"image"
	"testing"

	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_snap"
	"github.com/stretchr/testify/assert"
)

func TestPalettedImage(t *testing.T) {
	rendura.SetScreenSize(2, 3)
	rendura.Palette[1] = 0xffaa44
	rendura.Palette[2] = 0xff0000
	rendura.Palette[3] = 0x00ff00
	rendura.Palette[4] = 0x0000ff
	rendura.Palette[5] = 0x00ffff
	rendura.Palette[6] = 0xff00ff
	screen := rendura.Screen()
	screen.SetAll(1, 2, 3, 4, 5, 6)
	// when
	img := rendura_snap.PalettedImage()
	// then
	assertPalettedImage(t, img, screen)
}

func assertPalettedImage(t *testing.T, img image.PalettedImage, screen rendura.Canvas) {
	t.Helper()

	// then color indexes are the same
	for y := 0; y < screen.H(); y++ {
		for x := 0; x < screen.W(); x++ {
			actual := img.ColorIndexAt(x, y)
			expected := screen.Get(x, y)
			assert.Equal(t, expected, actual)
		}
	}
	// and RGBA colors match
	for y := 0; y < screen.H(); y++ {
		for x := 0; x < screen.W(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			assert.Equal(t, uint8(0xff), uint8(a))
			actual := rendura.FromRGB(uint8(r), uint8(g), uint8(b))
			expected := rendura.Palette[screen.Get(x, y)]
			assert.Equal(t, expected, actual)
		}
	}
	// and size is the same
	assert.Equal(t,
		image.Rectangle{
			Max: image.Point{X: screen.W(), Y: screen.H()},
		},
		img.Bounds(),
		"image size is not same as screen size",
	)
}
