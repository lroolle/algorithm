package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
			if got := BubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
