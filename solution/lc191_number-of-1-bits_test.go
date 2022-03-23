package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	assert.Equal(t, 1, hammingWeight(1))
	assert.Equal(t, 1, hammingWeight(2))
	assert.Equal(t, 2, hammingWeight(3))
}

func Test_hammingWeightCleanLastBit(t *testing.T) {
	assert.Equal(t, 1, hammingWeightCleanLastBit(1))
	assert.Equal(t, 1, hammingWeightCleanLastBit(2))
	assert.Equal(t, 2, hammingWeightCleanLastBit(3))
}
