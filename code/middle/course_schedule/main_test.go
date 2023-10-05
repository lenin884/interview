package course_schedule

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanFinish(t *testing.T) {
	t.Run("should return true if there are no cycles", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}}

		actual := canFinish(numCourses, prerequisites)

		require.Equal(t, true, actual)
	})

	t.Run("should return false if there are cycles", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}, {0, 1}}

		actual := canFinish(numCourses, prerequisites)

		require.Equal(t, false, actual)
	})

	t.Run("should return true if there are no cycles", func(t *testing.T) {
		numCourses := 5
		prerequisites := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}

		actual := canFinish(numCourses, prerequisites)

		require.Equal(t, true, actual)
	})
}
