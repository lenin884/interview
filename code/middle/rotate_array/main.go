package rotate_array

//func rotate(nums []int, k int) {
//	// edge case
//	if len(nums) < 2 {
//		return
//	}
//
//	// k might be larger than len(nums)
//	k = k % len(nums)
//
//	// reverse the whole array
//	reverse(nums, 0, len(nums)-1)
//
//	// reverse the first k elements
//	reverse(nums, 0, k-1)
//
//	// reverse the last n-k elements
//	reverse(nums, k, len(nums)-1)
//}
//
//func reverse(nums []int, start, end int) {
//	for start < end {
//		// swap
//		nums[start], nums[end] = nums[end], nums[start]
//		start++
//		end--
//	}
//}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}

	count := 0 // Количество повернутых элементов
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}
