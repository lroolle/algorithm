package solution

import "testing"

var example1 = &TreeNode{
	Val:   5,
	Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
	Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
}

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{example1, 22}, true},
		{"Example 2", args{example1, 21}, false},
		{"Example 3", args{example1, 18}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum2(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum2() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum3(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum3() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum4(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
