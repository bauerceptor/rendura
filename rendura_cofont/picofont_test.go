package rendura_cofont_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_test_helpers"
)

//go:embed "font.png"
var fontPNG []byte

func TestPrint(t *testing.T) {
	t.Run("should print each character", func(t *testing.T) {
		rendura.SetScreenSize(128, 128)

		rendura.Palette = rendura.DecodePalette(fontPNG)
		canvas := rendura.DecodeCanvas(fontPNG)

		var table strings.Builder

		// print narrow characters
		for i := 16; i < 128; i++ { // skip escape codes below 16 (such as LF)
			table.WriteRune(rune(i))
			table.WriteByte(' ')
			if i%16 == 15 {
				table.WriteByte('\n')
			}
		}
		// print wide characters
		for i := 128; i < 256; i++ {
			table.WriteRune(rune(i))
			if i%16 == 15 {
				table.WriteByte('\n')
			}
		}
		rendura.SetColor(1)
		// when
		rendura_cofont.Print(table.String(), 0, 8)
		// then
		rendura_test_helpers.AssertSurfaceEqual(t, canvas, rendura.Screen())
	})
}
