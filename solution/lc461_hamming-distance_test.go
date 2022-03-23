package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingDistance(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 2))
	assert.Equal(t, 1, hammingDistance(2, 3))
}
