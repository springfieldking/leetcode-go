package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	nums := []int{3, 5, 10, 9, 4}
	rets := []int{4, 5, 5, 8, 8}
	for index, n := range nums {
		assert.Equal(t, rets[index], kl.Add(n))
	}
}
