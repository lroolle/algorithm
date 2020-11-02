package solution

import (
	"testing"
)

var t1 = &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
var t2 = &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
	Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
var tt = &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 4}},
	Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}}

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{t1, t2}, tt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); got == tt.want {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
