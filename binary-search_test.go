package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "last element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 1000},
			want: 7,
		},
		{
			name: "non-existing element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 2},
			want: -1,
		},
		{
			name: "non-existing negative element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: -344},
			want: -1,
		},
		{
			name: "non-existing large value",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 30000},
			want: -1,
		},
		{
			name: "first element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 43},
			want: 0,
		},
		{
			name: "middle element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 100},
			want: 3,
		},
		{
			name: "pre-last element",
			args: args{arr: []int{43, 56, 93, 100, 245, 673, 933, 1000}, val: 933},
			want: 6,
		},
		{
			name: "empty array",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
