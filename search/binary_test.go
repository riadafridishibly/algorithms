package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		key   int
		array []int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{
			name: "Found case middle",
			args: args{
				key:   3,
				array: []int{1, 2, 3, 4, 5, 6},
			},
			wantIndex: 2,
		},
		{
			name: "Found case first",
			args: args{
				key:   1,
				array: []int{1, 2, 3, 4, 5, 6},
			},
			wantIndex: 0,
		},
		{
			name: "Found case last",
			args: args{
				key:   6,
				array: []int{1, 2, 3, 4, 5, 6},
			},
			wantIndex: 5,
		},
		{
			name: "Not found case before",
			args: args{
				key:   -1,
				array: []int{1, 2, 3, 4, 5, 6},
			},
			wantIndex: -1,
		},
		{
			name: "Not found case after",
			args: args{
				key:   7,
				array: []int{1, 2, 3, 4, 5, 6},
			},
			wantIndex: -1,
		},
		{
			name: "Special case 0",
			args: args{
				key:   42,
				array: []int{},
			},
			wantIndex: -1,
		},
		{
			name: "Special case 1",
			args: args{
				key:   42,
				array: []int{42},
			},
			wantIndex: 0,
		},
		{
			name: "Special case 2",
			args: args{
				key:   43,
				array: []int{42},
			},
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := BinarySearch(tt.args.key, tt.args.array); gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
