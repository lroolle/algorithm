package solution

import (
	"reflect"
	"testing"
)

func Test_hanota(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{A: []int{2, 1, 0}, B: []int{}, C: []int{}}, []int{2, 1, 0}},
		{"Example 2", args{A: []int{1, 0}, B: []int{}, C: []int{}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hanota(tt.args.A, tt.args.B, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hanota() = %v, want %v", got, tt.want)
			}
		})
	}
}
