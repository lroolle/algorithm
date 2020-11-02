package solution

import "github.com/lroolle/algorithm/datastructure/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = tree.BinaryNode

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		ans := 0
		if node == nil {
			return 0
		}
		if node.Left != nil {
			if isLeaf(node.Left) {
				ans += node.Left.Val
			} else {
				ans += dfs(node.Left)
			}
		}
		if node.Right != nil && !isLeaf(node.Right) {
			ans += dfs(node.Right)
		}
		return ans
	}

	return dfs(root)
}
