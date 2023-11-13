package candy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCandy(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equalf(t, 5, candy([]int{1, 0, 2}), "candy([]int{1, 0, 2}) should be 5")
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equalf(t, 4, candy([]int{1, 2, 2}), "candy([]int{1, 2, 2}) should be 4")
	})
}
