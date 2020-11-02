package solution

import (
	"reflect"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
