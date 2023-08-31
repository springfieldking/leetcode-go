package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	assert.Equal(t, 3, removeDuplicates([]int{1, 2, 2, 3, 3, 3}))
}
