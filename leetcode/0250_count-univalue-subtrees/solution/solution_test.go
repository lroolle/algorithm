package solution

import "testing"

func Test_countUnivalSubtrees(t *testing.T) {
	root := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}},
	}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{root}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnivalSubtrees(tt.args.root); got != tt.want {
				t.Errorf("countUnivalSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
