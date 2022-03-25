package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal(t *testing.T) {
	data := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	assert.Equal(t, 11, minimumTotal(data))
}
