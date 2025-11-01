/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	dummy := new(ListNode)
	curr := dummy

	for e := head; e != nil; e = e.Next {
		if !hashMap[e.Val] {
			curr.Next = e
			curr = curr.Next
		}
	}
	curr.Next = nil

	return dummy.Next
}
