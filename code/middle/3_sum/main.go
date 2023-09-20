package __sum

import "sort"

// threeSum function return three pairs when nums[i] + nums[j] + nums[k] = 0 and i != j != k
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	/*
		Implementation of the algorithm:
		1. Sort the input array. Sorting will help in efficiently finding the triplets that sum up to the target because it allows us to use two-pointer approach while iterating through the array.
		2. Iterate through the sorted array using a for loop, considering each element as a potential first element of the triplet.
		3. Inside the loop, use two pointers, one pointing to the element after the current element, and the other pointing to the end of the array.
		4. Calculate the current sum using the first element and the two pointers' elements.
		5. If the current sum is equal to the target sum, add the triplet (three elements) to the result list.
		6. If the current sum is less than the target, increment the left pointer.
		7. If the current sum is greater than the target, decrement the right pointer.
		8. Continue this process until both pointers meet.
		9. Ensure you skip duplicates to avoid duplicate triplets.
		10. Continue the loop until all elements have been considered as potential first elements.
	*/
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip duplicates
			continue
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { // skip duplicates
					left++
				}

				for left < right && nums[right] == nums[right-1] { // skip duplicates
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
