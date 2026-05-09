package rendura_font_test

import (
	_ "embed"
	"github.com/bauerceptor/rendura/rendura_test_helpers"
	"testing"

	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_font"
)

//go:embed internal/test/font.png
var fontPNG []byte

var fontSheet rendura_font.Sheet

func init() {
	prevPalette := rendura.Palette
	defer func() {
		rendura.Palette = prevPalette
	}()

	rendura.Palette = rendura.DecodePalette(fontPNG)
	fontCanvas := rendura.DecodeCanvas(fontPNG)
	fontSheet = rendura_font.Sheet{
		Chars: map[rune]rendura.Sprite{
			'S': {
				Area:   rendura.Area[int]{X: 0, Y: 0, W: 8, H: 8},
				Source: fontCanvas,
			},
			'T': {
				Area:   rendura.Area[int]{X: 8, Y: 0, W: 8, H: 8},
				Source: fontCanvas,
			},
			'⬤': {
				Area:   rendura.Area[int]{X: 0, Y: 8, W: 8, H: 8},
				Source: fontCanvas,
			},
			'❤': {
				Area:   rendura.Area[int]{X: 8, Y: 8, W: 8, H: 8},
				Source: fontCanvas,
			},
		},
		Height:  8,
		FgColor: 1,
		BgColor: 0,
	}
}

var (
	//go:embed internal/test/text-color-equal-to-bg.png
	textColorEqualToBg []byte
	//go:embed internal/test/text-color-equal-to-fg.png
	textColorEqualToFg []byte
	//go:embed internal/test/text-color-different.png
	textColorDifferent []byte
)

func TestSheet_Print(t *testing.T) {
	t.Run("should print with different colors", func(t *testing.T) {
		tests := map[string]struct {
			bgColor   rendura.Color
			textColor rendura.Color
			png       []byte
		}{
			"text color different than Bg and Fg": {
				bgColor:   0,
				textColor: 2,
				png:       textColorDifferent,
			},
			"text color equal to Bg": {
				bgColor:   2,
				textColor: 0,
				png:       textColorEqualToBg,
			},
			"text color equal to Fg": {
				bgColor:   0,
				textColor: 1,
				png:       textColorEqualToFg,
			},
		}

		for testName, testCase := range tests {
			t.Run(testName, func(t *testing.T) {
				rendura.Palette = rendura.DecodePalette(testCase.png)
				expectedCanvas := rendura.DecodeCanvas(testCase.png)
				rendura.SetScreenSize(8, 8)

				rendura.Screen().Clear(testCase.bgColor)
				rendura.SetColor(testCase.textColor)
				rendura.SetTransparency(testCase.textColor, false)
				// when
				fontSheet.Print("S", 0, 0)
				// then
				rendura_test_helpers.AssertSurfaceEqual(t, expectedCanvas, rendura.Screen())
			})
		}
	})
}

func BenchmarkSheet_Print(b *testing.B) {
	sheet := rendura_font.Sheet{
		Chars: map[rune]rendura.Sprite{
			'a': {
				Area:   rendura.Area[int]{X: 0, Y: 0, W: 8, H: 8},
				Source: rendura.NewCanvas(8, 8),
			},
		},
	}
	b.ReportAllocs()
	for b.Loop() {
		sheet.Print("aaaaaaaaaaaaaaaaaaaa", 100, 100)
	}
}
