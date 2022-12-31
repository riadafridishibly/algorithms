package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic merge sort",
			args: args{
				array: []int{9, 2, 3, 1, 8, 7, 5, 4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.array)
			if !isSorted(tt.args.array) {
				t.Errorf("%v: not sorted", tt.name)
			}
		})
	}
}
