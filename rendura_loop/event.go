package rendura_loop

type Event string

const (
	EventInit        Event = "init"         // when the game is started, just before the first frame
	EventFrameStart  Event = "frame_start"  // beginning of the frame
	EventUpdate      Event = "update"       // after rendura.Update
	EventLateUpdate  Event = "late_update"  // after EventUpdate
	EventDraw        Event = "draw"         // after rendura.Draw
	EventLateDraw    Event = "late_draw"    // after EventDraw
	EventWindowClose Event = "window_close" // when a user closes the window (desktop only)
)
