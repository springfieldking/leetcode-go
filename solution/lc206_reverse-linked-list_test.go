package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	// create list
	arr := []int{1, 2, 3, 4, 5}
	var head *ListNode = nil
	for i := len(arr) - 1; i >= 0; i-- {
		node := &ListNode{Val: arr[i], Next: head}
		head = node
	}

	// do reverse
	newHead := reverseList(head)

	// check
	for index, _ := range arr {
		node := newHead
		for i := 0; i < index; i++ {
			node = node.Next
		}
		assert.Equal(t, arr[len(arr)-1-index], node.Val)
	}

}
