/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)

	back, prev, curr := dummy, head, head.Next
	for curr != nil {
		prev.Next = curr.Next
		curr.Next = prev
		back.Next = curr

		back = prev
		prev = prev.Next
		if prev == nil {
			break
		}
		curr = prev.Next
	}

	return dummy.Next
}
