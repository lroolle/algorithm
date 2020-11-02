package solution

import "github.com/lroolle/algorithm/datastructure/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.BinaryNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	cur := root
	for {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
}
