package solution

import (
	"github.com/lroolle/algorithm/datastructure/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.BinaryNode

func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, pathSum int) {
		if node == nil {
			return
		}
		pathSum -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && pathSum == 0 {
			ret = append(ret, append([]int{}, path...))
		}
		dfs(node.Left, pathSum)
		dfs(node.Right, pathSum)
		path = path[:len(path)-1]
	}

	dfs(root, sum)
	return ret
}
