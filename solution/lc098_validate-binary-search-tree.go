package solution

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(curr *TreeNode, lower, upper int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= lower || curr.Val >= upper {
		return false
	}

	return isValidBSTHelper(curr.Left, lower, curr.Val) && isValidBSTHelper(curr.Right, curr.Val, upper)
}
