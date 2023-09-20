package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	hashmap := make(map[int]bool)

	for _, num := range nums {
		hashmap[num] = true
	}

	longestStreak := 0

	for num := range hashmap {
		if !hashmap[num-1] {
			currentNum := num
			currentStreak := 1

			for hashmap[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
