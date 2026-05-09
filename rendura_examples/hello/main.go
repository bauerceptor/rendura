package main

import (
	"github.com/bauerceptor/rendura"                   // import pixelforge core package
	"github.com/bauerceptor/rendura/rendura_cofont" // import very small pico-8 font
	"github.com/bauerceptor/rendura/rendura_ebiten" // import backend
)

func main() {
	rendura.SetScreenSize(47, 9) // set custom screen size
	rendura.Draw = func() {      // draw will be executed each frame
		rendura_cofont.Print("HELLO WORLD", 2, 2)
	}
	rendura_ebiten.Run() // run backend
}
