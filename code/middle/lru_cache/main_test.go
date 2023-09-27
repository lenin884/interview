package lru_cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		obj := Constructor(2)
		obj.Put(1, 1)
		obj.Put(2, 2)
		require.Equal(t, 1, obj.Get(1))
		obj.Put(3, 3)
		require.Equal(t, -1, obj.Get(2))
		obj.Put(4, 4)
		require.Equal(t, -1, obj.Get(1))
		require.Equal(t, 3, obj.Get(3))
		require.Equal(t, 4, obj.Get(4))
	})
}
