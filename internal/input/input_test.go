package input_test

import (
	"github.com/bauerceptor/rendura/pkg/rendura"
	"github.com/bauerceptor/rendura/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState_Duration(t *testing.T) {
	const btn = "btn"

	t.Run("should return 0 when input was never pressed", func(t *testing.T) {
		var i input.State[string]
		assert.Equal(t, 0, i.Duration(btn))
	})

	t.Run("should return 1 when input was pressed and released this frame", func(t *testing.T) {
		var i input.State[string]
		rendura.Frame = 1
		i.SetDownFrame(btn, rendura.Frame)
		i.SetUpFrame(btn, rendura.Frame)
		assert.Equal(t, 1, i.Duration(btn))
	})

	t.Run("should return 0 when input was pressed previous frame and released this frame", func(t *testing.T) {
		var i input.State[string]
		rendura.Frame = 0
		i.SetDownFrame(btn, rendura.Frame)
		rendura.Frame++
		i.SetUpFrame(btn, rendura.Frame)
		assert.Equal(t, 0, i.Duration(btn))
	})

	t.Run("should return 2 when input was pressed previous frame but not released this frame", func(t *testing.T) {
		var i input.State[string]
		rendura.Frame = 0
		i.SetDownFrame(btn, rendura.Frame)
		rendura.Frame++
		assert.Equal(t, 2, i.Duration(btn))
	})

	t.Run("should return 1 when input was pressed, released and pressed again this frame", func(t *testing.T) {
		var i input.State[string]
		rendura.Frame = 0
		i.SetDownFrame(btn, rendura.Frame)
		i.SetUpFrame(btn, rendura.Frame)
		i.SetDownFrame(btn, rendura.Frame)
		assert.Equal(t, 1, i.Duration(btn))
	})

	t.Run("should return 1 when input was pressed and released previous frame and pressed this frame", func(t *testing.T) {
		var i input.State[string]
		rendura.Frame = 0
		i.SetDownFrame(btn, rendura.Frame)
		i.SetUpFrame(btn, rendura.Frame)
		rendura.Frame++
		i.SetDownFrame(btn, rendura.Frame)
		assert.Equal(t, 1, i.Duration(btn))
	})
}
