package rendura_mouse

import (
	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/internal/input"
	pievent "github.com/bauerceptor/rendura/rendura_event"
)

var (
	Position      rendura.Position
	MovementDelta rendura.Position // mouse movement delta since the last frame
)

func Duration(b Button) int {
	return buttonState.Duration(b)
}

type Button string

const (
	Left  Button = "Left"
	Right Button = "Right"
)

var buttonTarget = pievent.NewTarget[EventButton]()
var buttonDebugTarget = pievent.NewTarget[EventButton]()
var moveTarget = pievent.NewTarget[EventMove]()
var moveDebugTarget = pievent.NewTarget[EventMove]()

var buttonState input.State[Button]

func init() {
	onButton := func(event EventButton, _ pievent.Handler) {
		switch event.Type {
		case EventButtonDown:
			buttonState.SetDownFrame(event.Button, rendura.Frame)
		case EventButtonUp:
			buttonState.SetUpFrame(event.Button, rendura.Frame)
		}
	}
	buttonTarget.SubscribeAll(onButton)
	buttonDebugTarget.SubscribeAll(onButton)

	onMove := func(event EventMove, _ pievent.Handler) {
		Position = event.Position
		MovementDelta = event.Position.Subtract(event.Previous)
	}
	moveTarget.SubscribeAll(onMove)
	moveDebugTarget.SubscribeAll(onMove)
}
