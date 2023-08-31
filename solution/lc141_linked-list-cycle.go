package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// hasCycle dfs(或者bfs)+访问标记 判断上文是否碰到过
func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if _, ok := visited[node]; ok {
			return true
		} else {
			visited[node] = true
		}
		node = node.Next
	}
	return false
}
