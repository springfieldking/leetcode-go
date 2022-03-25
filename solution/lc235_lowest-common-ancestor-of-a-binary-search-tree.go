package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		// 左右子树都有，那么当前节点就是最最近的节点
		return root
	} else if left == nil {
		// 如果左边是空，返回右边
		return right
	} else {
		// 否则返回左边
		return left
	}
}
