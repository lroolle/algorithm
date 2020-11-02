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

func countUnivalSubtrees(root *TreeNode) int {
	var ret int

	var isUnival func(*TreeNode) bool
	isUnival = func(node *TreeNode) bool {
		// 1. This subtree does not have children
		if node.Left == nil && node.Right == nil {
			ret++
			return true
		}

		// 2. All the childre of this subtree own same value include this subtree itself
		isunival := true
		if node.Left != nil {
			isunival = isUnival(node.Left) && isunival && node.Left.Val == node.Val
		}
		if node.Right != nil {
			isunival = isUnival(node.Right) && isunival && node.Right.Val == node.Val
		}

		if isunival {
			ret++
		}
		return isunival
	}

	isUnival(root)
	return ret
}
