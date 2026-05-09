package rendura_key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_key"
)

func TestDuration(t *testing.T) {
	t.Run("should return 0 by default", func(t *testing.T) {
		assert.Equal(t, 0, rendura_key.Duration(rendura_key.A))
	})

	rendura_key.Target().Publish(aDownEvent)

	t.Run("should return 1 when key was down in the current frame", func(t *testing.T) {
		assert.Equal(t, 1, rendura_key.Duration(rendura_key.A))
	})

	rendura.Frame++

	t.Run("should return 2 when key has been down since the previous frame", func(t *testing.T) {
		assert.Equal(t, 2, rendura_key.Duration(rendura_key.A))
	})

	rendura_key.Target().Publish(aUpEvent)

	t.Run("should return 0 when key is up", func(t *testing.T) {
		assert.Equal(t, 0, rendura_key.Duration(rendura_key.A))
	})
}
