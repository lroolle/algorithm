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

func kthSmallest(root *TreeNode, k int) int {
	var ret []int
	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ret = append(ret, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return ret[k-1]
}

func kthSmallestStack(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
