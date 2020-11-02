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

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

// DFS, substract from sum
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if isLeaf(root) {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// DFS, substract from sum and backtracking
func hasPathSum2(root *TreeNode, sum int) bool {
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, remainSum int) bool {
		if node == nil {
			return false
		}
		if isLeaf(node) && remainSum == node.Val {
			return true
		}
		if isLeaf(node) {
			return false
		}
		if node.Left != nil && dfs(node.Left, remainSum-node.Val) {
			return true
		}
		if node.Right != nil && dfs(node.Right, remainSum-node.Val) {
			return true
		}
		return false
	}

	return dfs(root, sum)
}

type Item struct {
	node *TreeNode
	sum  int
}

// Stack, add to sum
func hasPathSum3(root *TreeNode, sum int) bool {
	stack := []Item{{node: root, sum: root.Val}}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		node, pathSum := cur.node, cur.sum
		stack = stack[:len(stack)-1]
		if isLeaf(node) && pathSum == sum {
			return true
		}
		if node.Left != nil {
			stack = append(stack, Item{node: node.Left, sum: pathSum + node.Left.Val})
		}
		if node.Right != nil {
			stack = append(stack, Item{node: node.Right, sum: pathSum + node.Right.Val})
		}
	}
	return false
}

// BFS, add to sum
func hasPathSum4(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := []Item{{root, root.Val}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node.Left == nil && cur.node.Right == nil {
			if cur.sum == sum {
				return true
			}
			continue
		}
		if cur.node.Left != nil {
			queue = append(queue, Item{cur.node.Left, cur.node.Left.Val + cur.sum})
		}
		if cur.node.Right != nil {
			queue = append(queue, Item{cur.node.Right, cur.node.Right.Val + cur.sum})
		}
	}
	return false
}
