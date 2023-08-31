package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	rightSideViewDfs(root, 0, &res)
	return res
}

func rightSideViewDfs(node *TreeNode, deep int, res *[]int) {
	if node == nil {
		return
	}
	if len(*res) == deep {
		*res = append(*res, node.Val)
	}
	deep++
	rightSideViewDfs(node.Right, deep, res)
	rightSideViewDfs(node.Left, deep, res)
}
