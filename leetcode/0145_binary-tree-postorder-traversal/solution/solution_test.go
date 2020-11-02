package solution

import (
	"reflect"
	"testing"
)

var example1 = &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
var want1 = []int{3, 2, 1}
var example2 *TreeNode
var want2 = []int(nil)
var example3 = &TreeNode{Val: 1}
var want3 = []int{1}
var example4 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
var want4 = []int{2, 1}
var example5 = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
var want5 = []int{2, 1}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{example1}, want1},
		{"Example 2", args{example2}, want2},
		{"Example 3", args{example3}, want3},
		{"Example 4", args{example4}, want4},
		{"Example 5", args{example5}, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %#v, want %#v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal2() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
