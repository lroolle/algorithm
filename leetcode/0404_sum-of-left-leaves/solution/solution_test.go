package solution

import "testing"

var example1 = &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{
	Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{example1}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
