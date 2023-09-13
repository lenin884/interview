package best_time_to_buy_and_sell_stock_2

func maxProfit(prices []int) int {
	/*
		Example: [7,1,5,3,6,4]
		Step 1: [0, 0, 4, 2, 5, 3]
	*/

	// make slice of profits
	profits := make([]int, len(prices))

	// calculate profits day by day
	for i := 1; i < len(prices); i++ {
		profits[i] = prices[i] - prices[i-1]
	}

	maxProfit := 0

	// sum all positive profits
	for _, profit := range profits {
		if profit > 0 {
			maxProfit += profit
		}
	}

	return maxProfit
}
