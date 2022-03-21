package solution

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 8, climbStairs(5))
}

func TestClimbStairsDP(t *testing.T) {
	assert.Equal(t, 8, climbStairs(5))
}

func TestClimbStairsDPSimplify(t *testing.T) {
	assert.Equal(t, 8, climbStairs(5))
}

func BenchmarkClimbStairs(b *testing.B) {
	fmt.Sprintln(climbStairsDP(1))
}
