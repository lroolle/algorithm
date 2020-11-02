package solution

import (
	"github.com/lroolle/algorithm/datastructure/list"
	reverse "github.com/lroolle/algorithm/leetcode/0206_reverse-linked-list/solution"
)

type ListNode = list.SingleListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	var hair = &ListNode{Next: head}
	var prev = hair

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		// next: 1->2-> [3]
		next := tail.Next
		// tail: prev->1-> [2] ->null, next point to null for reverse until tail
		tail.Next = nil
		head = reverse.ReverseList(head)

		// prev-> [1]
		tail = prev.Next
		// [0] ->2->1
		prev.Next = head
		// 2->1-> [3]
		tail.Next = next

		// 2-> [1] ->3->4
		prev = tail
		//1-> [3] ->4->
		head = next
	}
	return hair.Next
}
