package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange(t *testing.T) {
	cions := []int{1, 2, 5}
	assert.Equal(t, 1, coinChange(cions, 1))
	assert.Equal(t, 1, coinChange(cions, 2))
	assert.Equal(t, 2, coinChange(cions, 3))
}
