package solution

import (
	"reflect"
	"testing"
)

func Test_maxRectangle(t *testing.T) {
	type args struct {
		words []string
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
			if got := maxRectangle(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
