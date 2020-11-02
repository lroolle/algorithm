package solution

import (
	"github.com/lroolle/algorithm/datastructure/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = list.SingleListNode

func insertionSortList(head *ListNode) *ListNode {
	prev := &ListNode{}
	for head != nil {
		curr := prev
		next := head.Next
		for curr.Next != nil && head.Val.(int) > curr.Next.Val.(int) {
			curr = curr.Next
		}
		head.Next = curr.Next
		curr.Next = head
		head = next
	}
	return prev.Next
}
