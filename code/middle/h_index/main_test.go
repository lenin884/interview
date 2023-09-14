package h_index

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("5 elements got 3", func(t *testing.T) {
		citations := []int{3, 0, 6, 1, 5}
		require.Equal(t, hIndex(citations), 3)
	})

	t.Run("3 elements got 1", func(t *testing.T) {
		citations := []int{1, 3, 1}
		require.Equal(t, 1, hIndex(citations))
	})
}
