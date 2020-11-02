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

// Recursive
func postorderTraversal(root *TreeNode) (ret []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ret = append(ret, node.Val)
	}

	dfs(root)
	return
}

// Stack, traverse
func postorderTraversal2(root *TreeNode) (ret []int) {
	var prev *TreeNode
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ret = append(ret, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}
