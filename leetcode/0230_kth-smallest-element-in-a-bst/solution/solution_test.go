package solution

import "testing"

var root1 = &TreeNode{
	Val:   3,
	Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	Right: &TreeNode{Val: 4},
}

var root2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 4}},
	Right: &TreeNode{Val: 6},
}

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{root: root1, k: 1}, 1},
		{"Example 1", args{root: root2, k: 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestStack(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
