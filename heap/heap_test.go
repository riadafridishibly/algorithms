package heap_test

import (
	"testing"

	"github.com/riadafridishibly/algorithms/heap"
	"github.com/stretchr/testify/require"
)

func TestHeapify(t *testing.T) {
	a := []int{7, 1, 9, 2, 5, 3, 3, 4}
	heap.Heapify(a)
	require.Equal(t, 9, a[0])
}

func TestHeapSort(t *testing.T) {
	a := []int{7, 1, 9, 2, 5, 3, 3, 4}
	heap.HeapSort(a)
	require.IsNonDecreasing(t, a)
}

func TestMaxHeap(t *testing.T) {
	mh := heap.NewMaxHeap()
	a := []int{7, 1, 9, 2, 5, 3, 3, 4}
	for _, v := range a {
		mh.Add(v)
	}

	var res []int
	for mh.Len() > 0 {
		res = append(res, mh.PopMax())
	}
	require.IsNonIncreasing(t, res, "expected sorted in reverse order")
}
