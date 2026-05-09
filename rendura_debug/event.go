package rendura_debug

import pievent "github.com/bauerceptor/rendura/rendura_event"

type Event string

const (
	EventPause  Event = "pause"
	EventResume Event = "resume"
)

func Target() pievent.Target[Event] {
	return target
}

var target = pievent.NewTarget[Event]()
