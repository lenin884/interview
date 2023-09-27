package insert_interval

import "testing"

func TestName(t *testing.T) {
	t.Run("should return [[1,5],[6,9]] when intervals = [[1,3],[6,9]], newInterval = [2,5]", func(t *testing.T) {
		got := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
		want := [][]int{{1, 5}, {6, 9}}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

		for i := 0; i < len(got); i++ {
			if got[i][0] != want[i][0] || got[i][1] != want[i][1] {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})

	t.Run("should return [[1,2],[3,10],[12,16]] when intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]", func(t *testing.T) {
		got := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
		want := [][]int{{1, 2}, {3, 10}, {12, 16}}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

		for i := 0; i < len(got); i++ {
			if got[i][0] != want[i][0] || got[i][1] != want[i][1] {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})
}
