package solution

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
	return
}

func (kl *KthLargest) Pop() interface{} {
	len := kl.Len()
	v := kl.IntSlice[len-1]
	kl.IntSlice = kl.IntSlice[0 : len-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, n := range nums {
		kl.Add(n)
	}
	return kl
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
