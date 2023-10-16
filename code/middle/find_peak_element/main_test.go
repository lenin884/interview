package find_peak_element

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findPeak(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		require.Equal(t, 2, findPeakElement(nums))
	})
}
