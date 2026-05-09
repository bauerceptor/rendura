// This example demonstrates building a simple GUI hierarchy with rendura_gui.
// It shows:
//   - A panel (container) with a local coordinate system
//   - Three buttons arranged vertically inside the panel
//   - Clicking a button logs its label
//
// The layout is recalculated relative to the panel's position,
// showing rendura_gui's tree-structure approach.
package main

import (
	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_gui"
	"log"
)

// colors used in this example (default rendura palette):
const (
	lightBlue = 28
	white     = 7
	darkBlue  = 1
	lightGray = 6
	blue      = 12
)

func main() {
	rendura.SetScreenSize(128, 128)
	// create the root of the entire GUI element tree
	root := rendura_gui.New()
	// add a panel (container) at global coordinates
	panel := attachPanel(root, 32, 32, 63, 63)
	// add buttons to the panel using its local coordinate system
	attachButton(panel, 10, 9, 44, 14, "BUTTON 1")
	attachButton(panel, 10, 25, 44, 14, "BUTTON 2")
	// add a button with a callback that runs when the user clicks
	// and releases the left mouse button while staying inside its area
	btn3 := attachButton(panel, 10, 41, 44, 14, "BUTTON 3")
	btn3.OnTap = func(event rendura_gui.Event) {
		log.Println("Button 3 was tapped")
	}

	rendura.Update = func() {
		// root.Update() must be called in the game loop
		root.Update()
	}

	rendura.Draw = func() {
		rendura.Cls()
		// root.Draw() must be called in the game loop
		root.Draw()
	}

	rendura_ebiten.Run()
}

func attachPanel(parent *rendura_gui.Element, x, y, w, h int) *rendura_gui.Element {
	panel := rendura_gui.Attach(parent, x, y, w, h)
	panel.OnDraw = func(event rendura_gui.DrawEvent) {
		rendura.SetColor(lightBlue)
		rendura.Rect(0, 0, panel.W-1, panel.H-1)
		rendura.SetColor(darkBlue)
		rendura.RectFill(1, 1, panel.W-2, panel.H-2)
	}
	return panel
}

func attachButton(parent *rendura_gui.Element, x, y, w, h int, label string) *rendura_gui.Element {
	btn := rendura_gui.Attach(parent, x, y, w, h)
	btn.OnDraw = func(event rendura_gui.DrawEvent) {
		var frame, bg, text rendura.Color = lightGray, blue, white
		if event.HasPointer {
			frame, bg, text = lightGray, lightBlue, white
		}

		if event.Pressed {
			rendura.Camera.Y -= 1 // the camera is automatically reset after drawing the element
			bg = blue
		}

		rendura.SetColor(frame)
		rendura.Rect(0, 0, w-2, h-2)

		rendura.SetColor(bg)
		rendura.RectFill(1, 1, w-3, h-3)

		rendura.SetColor(text)
		rendura_cofont.Print(label, 6, 4)
	}
	return btn
}
