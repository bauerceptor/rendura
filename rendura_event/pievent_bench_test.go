package rendura_event_test

import (
	"testing"

	"github.com/bauerceptor/rendura/rendura_event"
)

func init() {
	rendura_event.GlobalTracingOff = true
}

func BenchmarkPublish(b *testing.B) {
	b.ReportAllocs()
	target := rendura_event.NewTarget[someEvent]()
	event := someEvent{a: "1"} // event should not be empty for get meaningful results
	target.SubscribeAll(func(someEvent, rendura_event.Handler) {})

	for b.Loop() {
		target.Publish(event) // zero alokacji! LOVE IT
	}
}

func BenchmarkSubscribe(b *testing.B) {
	b.ReportAllocs()
	target := rendura_event.NewTarget[someEvent]()
	listener := func(someEvent, rendura_event.Handler) {}

	for b.Loop() {
		// 3 allocs, because stack trace is analyzed - only for debugging
		// 0 allocs for production code
		target.SubscribeAll(listener)
	}
}

func BenchmarkSubscribeEvent(b *testing.B) {
	b.ReportAllocs()
	target := rendura_event.NewTarget[someEvent]()
	listener := func(someEvent, rendura_event.Handler) {}

	for b.Loop() {
		// 3 allocs, because stack trace is analyzed - only for debugging
		// 0 allocs for production code
		target.Subscribe(someEvent{a: "a"}, listener)
	}
}

type someEvent struct {
	a string
}
