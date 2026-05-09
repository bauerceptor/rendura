package rendura_pad_test

import (
	"testing"

	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_pad"
	"github.com/stretchr/testify/assert"
)

var (
	player0connected    = rendura_pad.EventConnection{Type: rendura_pad.EventConnect, Player: 0}
	player1connected    = rendura_pad.EventConnection{Type: rendura_pad.EventConnect, Player: 1}
	player0disconnected = rendura_pad.EventConnection{Type: rendura_pad.EventDisconnect, Player: 0}
	player1disconnected = rendura_pad.EventConnection{Type: rendura_pad.EventDisconnect, Player: 1}
)

func TestPlayerCount(t *testing.T) {
	assert.Equal(t, 0, rendura_pad.PlayerCount())
	rendura_pad.ConnectionTarget().Publish(player0connected)
	rendura_pad.ConnectionTarget().Publish(player1connected)
	assert.Equal(t, 2, rendura_pad.PlayerCount())
	rendura_pad.ConnectionTarget().Publish(player0disconnected)
	assert.Equal(t, 1, rendura_pad.PlayerCount())
	rendura_pad.ConnectionTarget().Publish(player1disconnected)
}

func TestDuration(t *testing.T) {
	{
		rendura_pad.ConnectionTarget().Publish(player0connected)
		rendura_pad.ButtonTarget().Publish(
			rendura_pad.EventButton{Type: rendura_pad.EventDown, Button: rendura_pad.A, Player: 0},
		)

		t.Run("should return duration when button was pressed", func(t *testing.T) {
			duration := rendura_pad.Duration(rendura_pad.A)
			assert.Equal(t, 1, duration)

			playerDuration := rendura_pad.PlayerDuration(rendura_pad.A, 0)
			assert.Equal(t, 1, playerDuration)
		})

		rendura.Frame++

		t.Run("should take into account how many frames passed", func(t *testing.T) {
			assert.Equal(t, 2, rendura_pad.Duration(rendura_pad.A))
			assert.Equal(t, 2, rendura_pad.PlayerDuration(rendura_pad.A, 0))
		})

		rendura_pad.ConnectionTarget().Publish(player0disconnected)
	}

	t.Run("should return 0 after controller was disconnected", func(t *testing.T) {
		assert.Equal(t, 0, rendura_pad.Duration(rendura_pad.A))
		assert.Equal(t, 0, rendura_pad.PlayerDuration(rendura_pad.A, 0))
	})

	t.Run("should return the longest duration when two controllers are pressed simultaneously", func(t *testing.T) {
		rendura_pad.ConnectionTarget().Publish(player0connected)
		defer rendura_pad.ConnectionTarget().Publish(player0disconnected)
		rendura_pad.ConnectionTarget().Publish(player1connected)
		defer rendura_pad.ConnectionTarget().Publish(player1disconnected)

		rendura_pad.ButtonTarget().Publish(
			rendura_pad.EventButton{Type: rendura_pad.EventDown, Button: rendura_pad.A, Player: 0},
		)
		rendura.Frame++
		rendura_pad.ButtonTarget().Publish(
			rendura_pad.EventButton{Type: rendura_pad.EventUp, Button: rendura_pad.A, Player: 1},
		)
		assert.Equal(t, 2, rendura_pad.Duration(rendura_pad.A))
		assert.Equal(t, 2, rendura_pad.PlayerDuration(rendura_pad.A, 0))
	})

	t.Run("should return 0 when player was never connected", func(t *testing.T) {
		assert.Equal(t, 0, rendura_pad.PlayerDuration(rendura_pad.A, 2))
	})
}
