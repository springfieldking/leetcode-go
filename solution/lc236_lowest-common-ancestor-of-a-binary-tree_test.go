package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor235(t *testing.T) {
	var root *TreeNode = &TreeNode{Val: 0}
	p := root
	q := root
	assert.Equal(t, root, lowestCommonAncestor235(root, p, q))
}
