package algorithms

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 numbers",
			args: args{arr: []int{94, 29, 343, 91, 2, 49, 1, -1, -4555, 8}},
			want: []int{-4555, -1, 1, 2, 8, 29, 49, 91, 94, 343},
		},
		{
			name: "array with one item",
			args: args{arr: []int{9}},
			want: []int{9},
		},
		{
			name: "7 numbers",
			args: args{arr: []int{-2, 53, 34, 2, 3, 0, 56}},
			want: []int{-2, 0, 2, 3, 34, 53, 56},
		},
		{
			name: "sorted numbers",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
