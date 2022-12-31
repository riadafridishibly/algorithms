package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic Insertion Sort",
			args: args{
				array: []int{8, 6, 4, 1, 7, 5, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.array)
			if !isSorted(tt.args.array) {
				t.Errorf("%v: Not sorted", tt.name)
			}
		})
	}
}
