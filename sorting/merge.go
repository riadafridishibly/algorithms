package sorting

import "github.com/riadafridishibly/algorithms/types"

// lo, hi inclusive
func merge[T types.Ordered](array []T, lo, mid, hi int) {
	i := lo
	j := mid + 1

	// Copy items
	copied := make([]T, len(array))
	copy(copied, array)

	for k := lo; k <= hi; k++ {
		if i > mid { // left half exhauseted
			array[k] = copied[j]
			j++
		} else if j > hi { // right half exhausted
			array[k] = copied[i]
			i++
		} else if copied[j] < copied[i] { // item on the right < item on the left
			array[k] = copied[j]
			j++
		} else { // item on the left <= item on the right
			array[k] = copied[i]
			i++
		}
	}
}

func MergeSortRange[T types.Ordered](array []T, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	MergeSortRange(array, lo, mid)
	MergeSortRange(array, mid+1, hi)
	merge(array, lo, mid, hi)
}

func MergeSort[T types.Ordered](array []T) {
	MergeSortRange(array, 0, len(array)-1)
}
