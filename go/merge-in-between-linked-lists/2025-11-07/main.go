/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	firstNode, secondNode := list1, list1

	length := 0
	for e := list1; e != nil; e = e.Next {
		if length+1 == a {
			firstNode = e
		} else if length-1 == b {
			secondNode = e
			break
		}
		length++
	}

	firstNode.Next = list2
	current := list2
	for current.Next != nil {
		current = current.Next
	}
	current.Next = secondNode

	return list1
}
