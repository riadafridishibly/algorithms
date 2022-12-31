package sorting

import (
	"testing"

	"github.com/riadafridishibly/algorithms/types"
)

func isSorted[T types.Ordered](values []T) bool {
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			return false
		}
	}
	return true
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic sort",
			args: args{
				array: []int{9, 1, 3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.array)
			if !isSorted(tt.args.array) {
				t.Errorf("%v: array not sorted", tt.name)
			}
		})
	}
}
