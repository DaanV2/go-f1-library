package f1_2023_test

import (
	"testing"

	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
	"github.com/stretchr/testify/require"
)

func Test_Game_Content(t *testing.T) {
	t.Run("Drivers", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Drivers)
	})
	t.Run("Nationalities", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Nationalities)
	})
	t.Run("Teams", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Teams)
	})
	t.Run("Tracks", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Tracks)
	})
	t.Run("Penalties", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Penalties)
	})
	t.Run("Infringements", func(t *testing.T) {
		testCollection(t, f1_2023.Game.Infringements)
	})
}

type idItem interface {
	Id() uint8
	String() string
}

type collection[T idItem] interface {
	Get(id uint8) T
	Max() uint8
}

func testCollection[T collection[U], U idItem](t *testing.T, c T) {
	max := c.Max()

	for i := uint8(0); i <= max; i++ {
		item := c.Get(i)

		id := item.Id()
		require.Equal(t, i, id)

		str := item.String()
		require.NotEmpty(t, str)
	}
}
