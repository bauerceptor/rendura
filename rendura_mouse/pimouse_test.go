package rendura_mouse_test

import (
	"testing"

	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/rendura_mouse"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Run("should return 0 by default", func(t *testing.T) {
		assert.Equal(t, 0, rendura_mouse.Duration(rendura_mouse.Left))
	})

	rendura_mouse.ButtonTarget().Publish(rendura_mouse.EventButton{
		Button: rendura_mouse.Left,
		Type:   rendura_mouse.EventButtonDown,
	})

	t.Run("should return 1 when button was down in the current frame", func(t *testing.T) {
		assert.Equal(t, 1, rendura_mouse.Duration(rendura_mouse.Left))
	})

	rendura.Frame++

	t.Run("should return 2 when button has been down since the previous frame", func(t *testing.T) {
		assert.Equal(t, 2, rendura_mouse.Duration(rendura_mouse.Left))
	})

	rendura_mouse.ButtonTarget().Publish(rendura_mouse.EventButton{
		Button: rendura_mouse.Left,
		Type:   rendura_mouse.EventButtonUp,
	})

	t.Run("should return 0 when button is up", func(t *testing.T) {
		assert.Equal(t, 0, rendura_mouse.Duration(rendura_mouse.Left))
	})
}

func TestPosition(t *testing.T) {
	expected := rendura.Position{X: 1, Y: 2}
	event := rendura_mouse.EventMove{
		Position: expected,
		Previous: rendura.Position{X: 3, Y: 5},
	}
	// when
	rendura_mouse.MoveTarget().Publish(event)
	// then
	assert.Equal(t, expected, rendura_mouse.Position)
	assert.Equal(t, rendura.Position{X: -2, Y: -3}, rendura_mouse.MovementDelta)
}
