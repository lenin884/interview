package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}

	result = append(result, newInterval)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
