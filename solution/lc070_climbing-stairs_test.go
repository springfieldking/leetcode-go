package solution

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, 8, climbStairs(5))
	assert.Equal(t, 8, climbStairsDP(5))
	assert.Equal(t, 8, climbStairsDPSimplify(5))
}

func Benchmark_climbStairs(b *testing.B) {
	fmt.Sprintln(climbStairsDP(1))
}
