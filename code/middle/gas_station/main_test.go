package gas_station

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		gas := []int{1, 2, 3, 4, 5}
		cost := []int{3, 4, 5, 1, 2}

		expected := 3

		require.Equal(t, expected, canCompleteCircuit(gas, cost))
	})

	t.Run("return -1", func(t *testing.T) {
		gas := []int{2, 3, 4}
		cost := []int{3, 4, 3}

		expected := -1

		require.Equal(t, expected, canCompleteCircuit(gas, cost))
	})
}
