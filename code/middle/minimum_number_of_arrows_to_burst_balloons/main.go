package minimum_number_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// sort by starts
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	prev := points[0]
	total := 1

	for i := 0; i < len(points); i++ {
		start := points[i][0]
		end := points[i][1]
		if start > prev[1] {
			total++
			prev = points[i]
		} else {
			prev[1] = min(prev[1], end)
		}
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
