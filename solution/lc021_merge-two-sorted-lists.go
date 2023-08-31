package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var hair *ListNode = &ListNode{}
	var prev = hair
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			prev = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			prev = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		prev.Next = list1
	} else if list2 != nil {
		prev.Next = list2
	}

	return hair.Next
}
