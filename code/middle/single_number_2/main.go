package single_number_2

func singleNumber(nums []int) int {
	var ones int
	var twos int

	for i := 0; i < len(nums); i++ {
		number := nums[i]
		ones = (ones ^ number) &^ twos
		twos = (twos ^ number) &^ ones
	}

	return ones
}
