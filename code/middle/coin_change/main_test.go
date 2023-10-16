package coin_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_coinChange(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 11
		expect := 3
		require.Equal(t, expect, coinChange(coins, amount))
	})

	t.Run("case 2", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		expect := -1
		require.Equal(t, expect, coinChange(coins, amount))
	})
}
