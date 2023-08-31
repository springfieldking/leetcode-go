package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	// 如果都小于当前节点，去左子树寻找
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor235(root.Left, p, q)
	}
	// 如果都大雨当前节点，去右子树寻找
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor235(root.Right, p, q)
	}
	// 如果分叉，说明当前节点就是
	return root
}
