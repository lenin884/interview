package h_index

import "sort"

/*
hIndex time complexity O(nlogn)
space complexity O(1)
Функция которая возвращает максимальное число n, такое что у автора n статей, в которых n раз упоминается его авторство.
Алгоритм следующий:
1. Сортируем массив по возрастанию
2. Ищем первый индекс, где citations[i] >= n - i
*/
func hIndex(citations []int) int {
	// sort citations in descending order
	sort.Ints(citations)

	// find the first index where citations[i] >= n - i
	n := len(citations)
	for i := 0; i < n; i++ {
		if citations[i] >= n-i {
			return n - i
		}
	}

	return 0
}
