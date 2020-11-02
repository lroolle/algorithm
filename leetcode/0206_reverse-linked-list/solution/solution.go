package solution

import "github.com/lroolle/algorithm/datastructure/list"

type ListNode = list.SingleListNode

// Iterative
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// Recursive
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var curr = reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}

func ReverseList(head *ListNode) *ListNode {
	return reverseList(head)
}
