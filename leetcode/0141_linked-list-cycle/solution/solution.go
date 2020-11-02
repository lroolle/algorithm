package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var p1, p2 *ListNode
	p1, p2 = head.Next, head

	for p1 != p2 {
		if p1 == nil || p1.Next == nil {
			return false
		}
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return true
}
