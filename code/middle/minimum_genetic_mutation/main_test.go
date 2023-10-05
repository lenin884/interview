package minimum_genetic_mutation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinMutation(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		startGene := "AACCGGTT"
		endGene := "AACCGGTA"
		bank := []string{"AACCGGTA"}

		actual := minMutation(startGene, endGene, bank)

		require.Equal(t, 1, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		startGene := "AACCGGTT"
		endGene := "AAACGGTA"
		bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

		actual := minMutation(startGene, endGene, bank)

		require.Equal(t, 2, actual)
	})
}
