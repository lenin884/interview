package candy

func candy(ratings []int) int {
	cnd := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		cnd[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cnd[i] = cnd[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			cnd[i] = max(cnd[i], cnd[i+1]+1)
		}
	}

	sum := 0
	for _, v := range cnd {
		sum += v
	}

	return sum
}
