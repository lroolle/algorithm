package solution

import (
	"reflect"
	"testing"
)

var example1 = &TreeNode{
	Val:   5,
	Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
	Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
}

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{example1, 22}, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{"Example 2", args{example1, 21}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
