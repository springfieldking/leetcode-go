package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	var find = twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, len(find), 2)
	assert.Equal(t, find[0], 0)
	assert.Equal(t, find[1], 1)
}
