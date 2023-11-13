package trapping_rain_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrap(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equalf(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}) should be 6")
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equalf(t, 9, trap([]int{4, 2, 0, 3, 2, 5}), "trap([]int{4,2,0,3,2,5}) should be 9")
	})
}
