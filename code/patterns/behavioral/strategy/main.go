package main

/*
Давайте рассмотрим пример использования паттерна "Стратегия" для реализации сортировки массива
с использованием различных алгоритмов сортировки.
*/

import "fmt"

// SortStrategy Интерфейс для стратегии сортировки
type SortStrategy interface {
	Sort([]int)
}

// BubbleSort Конкретная стратегия: сортировка пузырьком
type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// MergeSort Конкретная стратегия: сортировка слиянием
type MergeSort struct{}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func (ms *MergeSort) Sort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	mid := n / 2
	left := arr[:mid]
	right := arr[mid:]
	ms.Sort(left)
	ms.Sort(right)
	copy(arr, merge(left, right))
}

// Sorter Контекст
type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(arr []int) {
	s.strategy.Sort(arr)
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}

	bubbleSort := &BubbleSort{}
	mergeSort := &MergeSort{}

	sorter := NewSorter(bubbleSort)
	fmt.Println("Сортировка пузырьком:")
	sorter.Sort(arr)
	fmt.Println(arr)

	sorter.SetStrategy(mergeSort)
	fmt.Println("Сортировка слиянием:")
	sorter.Sort(arr)
	fmt.Println(arr)
}

/*
В этом примере мы определили две конкретные стратегии сортировки: сортировка пузырьком и сортировка слиянием.
Мы также создали контекст Sorter, который позволяет клиентам выбирать и устанавливать используемую стратегию сортировки.
*/
