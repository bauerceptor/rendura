package input

import (
	"github.com/bauerceptor/rendura"
	"github.com/hajimehoshi/ebiten/v2"
)

type Backend struct {
	Paused     *bool
	LeftOffset *float64
	TopOffset  *float64
	Scale      *float64

	keys          []ebiten.Key
	mousePosition rendura.Position
	gamepads      ebitenGamepads
}

func (g *Backend) Update() {
	g.updateMouse()
	g.updateKeyboard()
	g.gamepads.update()
}
