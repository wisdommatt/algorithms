package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 unordered numbers",
			args: args{arr: []int{93, 992, 34, 24, 92, 5, 34, 11, 20, 0}},
			want: []int{0, 5, 11, 20, 24, 34, 34, 92, 93, 992},
		},
		{
			name: "15 unordered numbers",
			args: args{arr: []int{
				3024, 6446, 2647, 5502, 9523,
				4432, 4140, 6344, 1766, 3384,
				3827, 1535, 8130, 2287, 5433,
			}},
			want: []int{
				1535, 1766, 2287, 2647, 3024,
				3384, 3827, 4140, 4432, 5433,
				5502, 6344, 6446, 8130, 9523,
			},
		},
		{
			name: "20 unordered numbers",
			args: args{arr: []int{
				1831, 2776, 9994, 9924, 7054,
				3804, 6032, 7104, 7734, 9671,
				3942, 9842, 3757, 2360, 8972,
				1979, 6145, 2799, 2164, 4108,
			}},
			want: []int{
				1831, 1979, 2164, 2360, 2776,
				2799, 3757, 3804, 3942, 4108,
				6032, 6145, 7054, 7104, 7734,
				8972, 9671, 9842, 9924, 9994,
			},
		},
		{
			name: "sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "1 item array",
			args: args{arr: []int{6}},
			want: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
