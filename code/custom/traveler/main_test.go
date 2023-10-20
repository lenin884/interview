package traveler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("find route", func(t *testing.T) {
		tickets := []Ticket{
			{"a", "c"},
			{"c", "d"},
			{"b", "e"},
			{"d", "b"},
			{"e", "f"},
		}
		expected := []string{"a", "c", "d", "b", "e", "f"}

		require.Equal(t, expected, findRoute(tickets))
	})

	t.Run("find route from alexey", func(t *testing.T) {
		tickets := []Ticket{
			{"a", "c"},
			{"c", "d"},
			{"b", "e"},
			{"d", "b"},
			{"e", "f"},
		}
		expected := []string{"a", "c", "d", "b", "e", "f"}

		require.Equal(t, expected, findRouteFromAlexey(tickets))
	})
}
