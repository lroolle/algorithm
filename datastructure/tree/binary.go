package tree

import "fmt"

type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

func (node *BinaryNode) Insert(val int) {
	if node == nil {
		return
	}
	if val < node.Val {
		if node.Left == nil {
			node.Left = &BinaryNode{Val: val}
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {

			node.Right = &BinaryNode{Val: val}
		} else {
			node.Right.Insert(val)
		}
	}
}

func (node *BinaryNode) PrintPreorder() {
	if node == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%#v ", node.Val)
	if node.Left != nil {
		node.Left.PrintPreorder()
	}
	if node.Right != nil {
		node.Right.PrintPreorder()
	}
}

func (node *BinaryNode) IterPrintPreorder() {
	stack := []*BinaryNode{node}
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		fmt.Printf("%#v ", node.Val)

		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func (node *BinaryNode) PrintInorder() {
	if node == nil {
		fmt.Println()
		return
	}
	if node.Left != nil {
		node.Left.PrintInorder()
	}
	fmt.Printf("%#v ", node.Val)
	if node.Right != nil {
		node.Right.PrintInorder()
	}
}

func (node *BinaryNode) IterPrintInorder() {
	stack := []*BinaryNode{node}
	cur := node
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%#v ", cur.Val)
			cur = cur.Right
		}
	}
}

// Morris Inorder Traverse
func (node *BinaryNode) MorrisInorder() {
	cur := node
	for cur != nil {
		if cur.Left == nil {
			fmt.Printf("%#v ", cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			fmt.Printf("%#v ", cur.Val)
			cur = cur.Right
		}
	}
}

func (node *BinaryNode) PrintPostorder() {
	if node == nil {
		fmt.Println()
		return
	}
	if node.Left != nil {
		node.Left.PrintPostorder()
	}
	if node.Right != nil {
		node.Right.PrintPostorder()
	}
	fmt.Printf("%#v ", node.Val)
}

func (node *BinaryNode) IterPrintPostorder() {
	stack := []*BinaryNode{}
	cur := node
	for cur != nil || len(stack) > 0 {

	}
}

// Compare two binary trees values, return true if all values equal.
func (root1 *BinaryNode) Compare(root2 *BinaryNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return root1.Left.Compare(root2.Left) && root1.Right.Compare(root2.Right)
}

// Build a BinaryTree from leetcode test case, for example [2, 1, 3, null, 4]
// For more: https://support.leetcode-cn.com/hc/kb/article/1194353/
func TestCaseBuilder(in []interface{}) (root *BinaryNode) {
	if len(in) == 0 {
		return
	}

	// Reassign nodes for eaxmple {node{1, true}, node{2, true}, node{0, false}}
	type node struct {
		val int
		// Is a null node
		null bool
	}
	var nodes = make([]node, len(in))
	for i := len(in) - 1; i >= 0; i-- {
		switch in[i].(type) {
		case int:
			nodes[len(in)-1-i] = node{in[i].(int), false}
		default:
			nodes[len(in)-1-i] = node{0, true}
		}
	}

	root = &BinaryNode{Val: nodes[len(nodes)-1].val}
	nodes = nodes[:len(nodes)-1]

	q := []*BinaryNode{root}
	for len(q) > 0 {
		// Pop current binary tree node
		cur := q[0]
		q = q[1:]
		// left node
		if len(nodes) > 0 {
			left := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			if !left.null {
				cur.Left = &BinaryNode{Val: left.val}
				q = append(q, cur.Left)
			}
		}
		// right node
		if len(nodes) > 0 {
			right := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			if !right.null {
				cur.Right = &BinaryNode{Val: right.val}
				q = append(q, cur.Right)
			}
		}
	}

	return
}

type BinaryTree struct {
	root *BinaryNode
}
