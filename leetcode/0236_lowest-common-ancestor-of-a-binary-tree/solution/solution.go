package solution

import "github.com/lroolle/algorithm/datastructure/tree"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode = tree.BinaryNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// // p && q must in each side of current root node.
	// if left != nil && right != nil {
	// 	return root
	// }
	// // non-nil right and nil left, p q in right side of current root node.
	// if right != nil {
	// 	return right
	// }
	// return left
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var ret *TreeNode
	var dfs func(*TreeNode, *TreeNode, *TreeNode) bool

	dfs = func(node, p, q *TreeNode) bool {
		if node == nil {
			return false
		}

		in_left := dfs(node.Left, p, q)
		in_right := dfs(node.Right, p, q)
		if (in_left && in_right) ||
			((node.Val == p.Val || node.Val == q.Val) && (in_left || in_right)) {
			ret = node
		}
		return in_left || in_right || (node.Val == p.Val || node.Val == q.Val)
	}

	dfs(root, p, q)
	return ret
}
