package list

import (
	"fmt"
	"strings"
)

type SingleListNode struct {
	Val  interface{}
	Next *SingleListNode
}

func (node *SingleListNode) Print() {
	var curr = node
	for curr != nil {
		fmt.Printf("%+v->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func (node *SingleListNode) Sprint() string {
	var sb strings.Builder
	var curr = node
	for curr != nil {
		fmt.Fprintf(&sb, "%+v->", curr.Val)
		curr = curr.Next
	}
	return sb.String()
}

func (node *SingleListNode) Insert(val interface{}) *SingleListNode {
	newNode := &SingleListNode{Val: val}
	if node == nil || node.Val == nil {
		return newNode
	}
	prev := &SingleListNode{Next: node}
	curr := prev
	for curr.Next != nil && curr.Next.Val.(int) < val.(int) {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode
	return prev.Next
}

func (node1 *SingleListNode) Compare(node2 *SingleListNode) bool {
	cur1, cur2 := node1, node2
	for cur1 != nil || cur2 != nil {
		if cur1 == nil || cur2 == nil {
			return false
		}
		if cur1.Val != cur2.Val {
			return false
		}
		cur1, cur2 = cur1.Next, cur2.Next
	}
	return true
}

type DoubleListNode struct {
	Prev *DoubleListNode
	Next *DoubleListNode
	Val  interface{}
}

func (node *DoubleListNode) Print() {
	var prev = node

	for prev.Prev != nil {
		prev = prev.Prev
	}

	var cur = prev
	for cur != nil {
		fmt.Printf("%+v<->", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
