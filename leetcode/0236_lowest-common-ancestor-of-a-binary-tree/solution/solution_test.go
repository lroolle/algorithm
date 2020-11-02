package solution

import (
	"reflect"
	"testing"
)

var p1 = &TreeNode{Val: 3}
var q1 = &TreeNode{Val: 5}
var ancestor1 = &TreeNode{Val: 4, Left: p1, Right: q1}
var example1 = &TreeNode{
	Val: 6,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 0},
		Right: ancestor1,
	}}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{example1, p1, q1}, ancestor1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}
