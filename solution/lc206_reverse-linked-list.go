package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head
	for curr != nil {
		var next *ListNode = curr.Next

		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
