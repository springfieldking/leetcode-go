package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var hair = &ListNode{Next: head}
	var groupPrev = hair
	var groupHead = head
	for groupHead != nil {
		// 找到kgroup的头部和尾部，如果不够直接返回
		var groupTail = groupPrev
		for i:=0; i<k; i++ {
			groupTail = groupTail.Next
			if groupTail == nil{
				return hair.Next
			}
		}

		// 记录group的next
		groupNext := groupTail.Next
		// 组内反转
		newGroupHead, newGroupTail := reverseInGroup(groupHead, groupTail)

		// 新的head和tail指针链接
		groupPrev.Next = newGroupHead
		newGroupTail.Next = groupNext

		// 跟新groupprev和grouphead，继续下一个循环
		groupPrev = newGroupTail
		groupHead = groupNext
	}
	return hair.Next
}

func reverseInGroup(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head
	for prev != tail {
		next := curr.Next

		curr.Next = prev
		prev = curr
		curr = next
	}
	return tail, head
}
