package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	assert.Equal(t, true, isValidBST(nil))
}
