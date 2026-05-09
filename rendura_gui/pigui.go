// Package pigui offers a minimal API for building GUIs.
package rendura_gui

import (
	"slices"

	"github.com/bauerceptor/rendura"
	pimouse "github.com/bauerceptor/rendura/rendura_mouse"
)

// New creates a new GUI root element with the current screen size.
//
// To update and draw the element along with its children,
// add it to your game loop by calling Element.Update and Element.Draw.
func New() *Element {
	return &Element{
		Area: rendura.IntArea{
			W: rendura.Screen().W(),
			H: rendura.Screen().H(),
		},
	}
}

// Attach attaches a new element with the specified size to the parent.
//
// It returns the newly created element.
func Attach(parent *Element, x, y, w, h int) *Element {
	ch := &Element{
		Area: rendura.IntArea{X: x, Y: y, W: w, H: h},
	}
	parent.Attach(ch)
	return ch
}

type Element struct {
	rendura.Area[int]

	OnDraw    func(DrawEvent)
	OnUpdate  func(UpdateEvent)
	OnPress   func(Event)
	OnRelease func(Event)
	OnTap     func(Event)
	children  []*Element
	pressed   bool
}

// Attach re-attaches an existing element to the parent e.
func (e *Element) Attach(child *Element) {
	e.children = append(e.children, child)
}

// Detach detaches the specified child element from e.
func (e *Element) Detach(child *Element) {
	e.children = slices.DeleteFunc(e.children, func(element *Element) bool {
		return child == element
	})
}

// Update should be called from within rendura.Update
// or by any subscriber listening to piloop.EventUpdate events
func (e *Element) Update() {
	prevCamera := rendura.Camera
	defer func() {
		rendura.Camera = prevCamera
	}()

	rendura.Camera.X -= e.X // I have to move the camera so that the children's children can pick up events
	rendura.Camera.Y -= e.Y

	mousePosition := pimouse.Position.Add(rendura.Camera)

	hasPointer := mousePosition.X >= 0 && mousePosition.X < e.W &&
		mousePosition.Y >= 0 && mousePosition.Y < e.H

	propagate := getPropagateToChildrenFromThePool()

	updateEvent := UpdateEvent{
		Element:                  e,
		HasPointer:               hasPointer,
		propagateToChildren:      propagate,
		propagateToChildrenToken: propagateToChildrenToken,
	}
	if e.OnUpdate != nil {
		e.OnUpdate(updateEvent)
	}

	mouseLeft := pimouse.Duration(pimouse.Left)

	if hasPointer && mouseLeft == 1 {
		e.pressed = true
		if e.OnPress != nil {
			e.OnPress(Event{
				Element:    e,
				HasPointer: true,
			})
		}
	} else if e.pressed && mouseLeft == 0 {
		e.pressed = false
		if e.OnRelease != nil {
			e.OnRelease(Event{
				Element:    e,
				HasPointer: hasPointer,
			})
		}
		if hasPointer {
			if e.OnTap != nil {
				e.OnTap(Event{
					Element:    e,
					HasPointer: true,
				})
			}
		}
	}

	childrenPropagation := propagate.value
	propagateToChildrenPool.Put(propagate)

	if childrenPropagation {
		for _, child := range e.children {
			child.Update()
		}
	}
}

// Draw should be called from within rendura.Draw
// or by any subscriber listening to piloop.EventDraw events
func (e *Element) Draw() {
	prevCamera := rendura.Camera
	defer func() {
		rendura.Camera = prevCamera
	}()

	rendura.Camera.X -= e.X
	rendura.Camera.Y -= e.Y

	prevClip := rendura.SetClip(rendura.IntArea{
		X: -rendura.Camera.X, Y: -rendura.Camera.Y,
		W: e.W, H: e.H,
	})
	defer func() {
		rendura.SetClip(prevClip)
	}()

	propagate := getPropagateToChildrenFromThePool()

	mousePosition := pimouse.Position.Add(rendura.Camera)
	hasPointer := mousePosition.X >= 0 && mousePosition.X < e.W &&
		mousePosition.Y >= 0 && mousePosition.Y < e.H

	drawEvent := DrawEvent{
		Element:                  e,
		HasPointer:               hasPointer,
		Pressed:                  e.pressed,
		propagateToChildren:      propagate,
		propagateToChildrenToken: propagateToChildrenToken,
	}
	if e.OnDraw != nil {
		e.OnDraw(drawEvent)
	}

	childrenPropagation := propagate.value
	propagateToChildrenPool.Put(propagate)

	if childrenPropagation {
		for _, child := range e.children {
			child.Draw()
		}
	}
}
