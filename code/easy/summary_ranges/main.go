package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := make([]string, 0)

	start := nums[0]

	/*
		Example: [0, 1, 2, 4, 5, 7]
		0 -> 1 -> 2 -> 4 -> 5 -> 7
		0 -> 2
		4 -> 5
		7
	*/
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d", start)+"->"+fmt.Sprintf("%d", nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[len(nums)-1] {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d", start)+"->"+fmt.Sprintf("%d", nums[len(nums)-1]))
	}

	return result
}
