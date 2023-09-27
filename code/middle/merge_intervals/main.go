package merge_intervals

import "sort"

//func merge(intervals [][]int) [][]int {
//	if len(intervals) == 0 {
//		return [][]int{}
//	}
//
//	result := make([][]int, 0)
//
//	firstPoint := intervals[0][0]
//	lastPoint := intervals[0][1]
//
//	for i := 1; i < len(intervals); i++ {
//		if intervals[i][0] <= lastPoint {
//			if intervals[i][1] > lastPoint {
//				lastPoint = intervals[i][1]
//			}
//		} else {
//			result = append(result, []int{firstPoint, lastPoint})
//			firstPoint = intervals[i][0]
//			lastPoint = intervals[i][1]
//		}
//	}
//
//	result = append(result, []int{firstPoint, lastPoint})
//
//	return result
//}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// sort first
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	result = append(result, []int{intervals[0][0], intervals[0][1]})

	// merge
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		lastEnd := result[len(result)-1][1]

		if start <= lastEnd {
			result[len(result)-1][1] = max(result[len(result)-1][1], end)
		} else {
			result = append(result, []int{start, end})
		}
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//func merge(intervals [][]int) [][]int {
//	if len(intervals) == 0 {
//		return nil
//	}
//	// sort
//	quickSort(intervals, 0, len(intervals)-1)
//	// merge
//	var res [][]int
//	res = append(res, intervals[0])
//	for i := 1; i < len(intervals); i++ {
//		if intervals[i][0] <= res[len(res)-1][1] {
//			if intervals[i][1] > res[len(res)-1][1] {
//				res[len(res)-1][1] = intervals[i][1]
//			}
//		} else {
//			res = append(res, intervals[i])
//		}
//	}
//	return res
//}
//
//func quickSort(intervals [][]int, left, right int) {
//	if left >= right {
//		return
//	}
//	pivot := partition(intervals, left, right)
//	quickSort(intervals, left, pivot-1)
//	quickSort(intervals, pivot+1, right)
//}
//
//func partition(intervals [][]int, left, right int) int {
//	pivot := right
//	i := left
//	for j := left; j < right; j++ {
//		if intervals[j][0] < intervals[pivot][0] {
//			intervals[i], intervals[j] = intervals[j], intervals[i]
//			i++
//		}
//	}
//	intervals[i], intervals[pivot] = intervals[pivot], intervals[i]
//	return i
//}

func quickSort(intervals []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partionalSort(intervals, left, right)
	quickSort(intervals, left, pivot-1)
	quickSort(intervals, pivot+1, right)
}

func partionalSort(intervals []int, left, right int) int {
	pivot := right
	i := left
	for j := left; j < right; j++ {
		if intervals[j] < intervals[pivot] {
			intervals[i], intervals[j] = intervals[j], intervals[i]
			i++
		}
	}
	intervals[i], intervals[pivot] = intervals[pivot], intervals[i]

	return i
}
