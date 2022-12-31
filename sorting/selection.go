package sorting

import (
	"github.com/riadafridishibly/algorithms/types"
)

func SelectionSort[T types.Ordered](array []T) {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}
