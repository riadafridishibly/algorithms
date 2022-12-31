package search

import (
	"github.com/riadafridishibly/algorithms/types"
)

func BinarySearch[T types.Ordered](key T, array []T) (index int) {
	lo := 0
	hi := len(array) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < array[mid] {
			hi = mid - 1
		} else if key > array[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
