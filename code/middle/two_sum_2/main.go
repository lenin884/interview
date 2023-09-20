package two_sum_2

func twoSum(numbers []int, target int) []int {
	sum, left, right := 0, 0, len(numbers)-1

	for left < right {
		sum = numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
