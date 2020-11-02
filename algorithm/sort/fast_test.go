package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"Example 2", args{[]int{1, 1, 1}}, []int{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sort(t *testing.T) {
	type args struct {
		nums []int
		l    int
		r    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.nums, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		nums []int
		l    int
		r    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.nums, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
