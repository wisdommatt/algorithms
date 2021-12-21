package algorithms

import "testing"

func TestKadaneMaxSumSubArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "array 1",
			args: args{arr: []int{3, 1, -4, 5, 9, -12}},
			want: 14,
		},
		{
			name: "array 2",
			args: args{arr: []int{-2, 0, 2, 3, -8, 8, 19, -10, 6}},
			want: 27,
		},
		{
			name: "array 2",
			args: args{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			want: 66,
		},
		{
			name: "empty array",
			args: args{arr: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KadaneMaxSumSubArray(tt.args.arr); got != tt.want {
				t.Errorf("KadaneMaxSumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
