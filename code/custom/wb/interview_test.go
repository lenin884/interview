package wb

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	t.Run("TestService", func(t *testing.T) {
		fmt.Println(predictable())
	})

	t.Run("ReadFromChan", func(t *testing.T) {
		t.Parallel()
		ReadFromChan()
	})

	t.Run("GetValue", func(t *testing.T) {
		fmt.Println(GetValue())
	})

	t.Run("cap", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		fmt.Println(cap(arr1))
		fmt.Println(cap(arr2))
		arr1 = append(arr1, arr2...)
		fmt.Println(cap(arr1))
	})
}
