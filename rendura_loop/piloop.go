// Package piloop defines events published during the game loop.
//
// It enables adding logic from any component,
// including those created by third parties.
package rendura_loop

import pievent "github.com/bauerceptor/rendura/rendura_event"

func Target() pievent.Target[Event] {
	return target
}

func DebugTarget() pievent.Target[Event] {
	return debugTarget
}

var (
	target      = pievent.NewTarget[Event]()
	debugTarget = pievent.NewTarget[Event]()
)
