package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if _, ok := visited[node]; ok {
			return node
		} else {
			visited[node] = true
		}
		node = node.Next
	}
	return nil
}