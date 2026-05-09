package rendura_key_test

import (
	"testing"

	"github.com/bauerceptor/rendura/rendura_key"
	"github.com/bauerceptor/rendura/rendura_loop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	aDownEvent    = rendura_key.Event{Type: rendura_key.EventDown, Key: rendura_key.A}
	aUpEvent      = rendura_key.Event{Type: rendura_key.EventUp, Key: rendura_key.A}
	ctrlDownEvent = rendura_key.Event{Type: rendura_key.EventDown, Key: rendura_key.Ctrl}
)

func TestRegisterShortcut(t *testing.T) {
	t.Run("single key", func(t *testing.T) {
		executionTimes := 0
		shortcut := rendura_key.RegisterShortcut(func() {
			executionTimes++
		}, rendura_key.A)
		require.NotNil(t, shortcut)

		rendura_key.Target().Publish(aDownEvent)
		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)

		t.Run("should execute callback when key is pressed", func(t *testing.T) {
			assert.Equal(t, 1, executionTimes)
		})

		executionTimes = 0

		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)

		t.Run("should not execute callback again on next frame when key is still pressed", func(t *testing.T) {
			assert.Equal(t, 0, executionTimes)
		})

		shortcut.Unregister()

		rendura_key.Target().Publish(aDownEvent)
		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)

		t.Run("should not execute callback after shortcut is unregistered", func(t *testing.T) {
			assert.Equal(t, 0, executionTimes)
		})
	})

	t.Run("multiple keys", func(t *testing.T) {
		executionTimes := 0
		shortcut := rendura_key.RegisterShortcut(func() {
			executionTimes++
		}, rendura_key.A, rendura_key.Ctrl)
		require.NotNil(t, shortcut)
		// when
		rendura_key.Target().Publish(aDownEvent)
		rendura_key.Target().Publish(ctrlDownEvent)
		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)
		// then
		assert.Equal(t, 1, executionTimes)
	})

	t.Run("should not run callback when all keys are not pressed simultaneously, but was down before", func(t *testing.T) {
		executionTimes := 0
		shortcut := rendura_key.RegisterShortcut(func() {
			executionTimes++
		}, rendura_key.A, rendura_key.Ctrl)
		require.NotNil(t, shortcut)

		rendura_key.Target().Publish(aDownEvent)
		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)
		require.Equal(t, 0, executionTimes)
		// when
		rendura_key.Target().Publish(ctrlDownEvent)
		rendura_key.Target().Publish(aUpEvent) // "A" no longer down
		rendura_loop.Target().Publish(rendura_loop.EventLateUpdate)
		// then
		assert.Equal(t, 0, executionTimes)
	})
}
