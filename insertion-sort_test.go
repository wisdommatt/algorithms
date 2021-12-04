package algorithms

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func BenchmarkInsetionSort(b *testing.B) {
	// benchmarking for sorting an array of 10k items.
	//
	// NOTE: running this test might take time on some machines
	// because of the number of operations required to perform
	// an insertion sort on an array.
	for i := 0; i < b.N; i++ {
		arr := []int{}
		for i := 10_000; i >= 0; i-- {
			arr = append(arr, i)
		}
		want := []int{}
		for i := 0; i <= 10_000; i++ {
			want = append(want, i)
		}
		InsertionSort(arr)
		if !reflect.DeepEqual(arr, want) {
			b.Errorf("InsertionSort() = %v, want %v", arr, want)
		}
	}
}
