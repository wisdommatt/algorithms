package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "array of 8 items",
			args: args{arr: []int{93, 42, 299, 19, 30, 12, 38, 82}},
			want: []int{12, 19, 30, 38, 42, 82, 93, 299},
		},
		{
			name: "array of 20 items",
			args: args{arr: []int{33, 9, 392, 239, 0, 49482, 89, 290, 23, 59, 20, 20, 95, 92, 59}},
			want: []int{0, 9, 20, 20, 23, 33, 59, 59, 89, 92, 95, 239, 290, 392, 49482},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
