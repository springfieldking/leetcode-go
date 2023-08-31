package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {

}

func Test_reverseSlice(t *testing.T) {
	// init
	head := &ListNode{1, nil}
	mid := &ListNode{2, nil}
	tail := &ListNode{3, nil}
	mid.Next = tail
	head.Next = mid
	assert.Equal(t, head.Val, 1)
	assert.Equal(t, mid.Val, 2)
	assert.Equal(t, tail.Val, 3)

	// reverse
	head, tail = reverseInGroup(head, tail)

	// 判断
	assert.Equal(t, head.Val, 3)
	assert.Equal(t, mid.Val, 2)
	assert.Equal(t, tail.Val, 1)
}
