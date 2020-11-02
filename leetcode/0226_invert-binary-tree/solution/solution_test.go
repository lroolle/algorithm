package solution

import (
	"reflect"
	"testing"
)

var root = &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}}

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{root}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.args.root)
			got.PrintPreorder()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
