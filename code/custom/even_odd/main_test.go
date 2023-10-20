package even_odd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvenBinary(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equalf(t, true, isEvenBinary(2), "isEvenBinary(2) should be true")
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equalf(t, false, isEvenBinary(33), "isEvenBinary(3) should be false")
	})
}

func TestEvenBinarySearch(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		require.Equalf(t, false, isEven2NSolution(37), "isEvenBinarySearch(2) should be true")
	})

	t.Run("case 2", func(t *testing.T) {
		require.Equalf(t, false, isEven2NSolution(33), "isEvenBinarySearch(3) should be false")
	})

	t.Run("case 3", func(t *testing.T) {
		require.Equalf(t, false, isEven2NSolution(579), "isEvenBinarySearch(2) should be true")
	})
}
