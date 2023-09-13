package best_time_to_buy_and_sell_stock_2

import "testing"

func TestName(t *testing.T) {

	t.Run("6 elements got 7", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}

		got := maxProfit(prices)
		want := 7

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("4 elements", func(t *testing.T) {
		prices := []int{4, 3, 2, 1}

		got := maxProfit(prices)
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
