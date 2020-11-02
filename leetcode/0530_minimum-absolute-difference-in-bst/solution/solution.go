package solution

import (
	"math"

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

func getMinimumDifference(root *TreeNode) int {
	ret, pre := math.MaxInt32, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ret {
			ret = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ret
}
