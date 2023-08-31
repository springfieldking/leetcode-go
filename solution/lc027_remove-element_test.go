package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{1, 2, 2, 3}, 2))
}
