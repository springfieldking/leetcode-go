package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	var root *TreeNode = nil
	assert.Equal(t, root, lowestCommonAncestor(nil, nil, nil))
}
