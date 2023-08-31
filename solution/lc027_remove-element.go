package solution

func removeElement(nums []int, val int) int {
	n := len(nums)
	left := 0
	right := 0
	for ; right < n; right++ {
		if nums[right] == val {
			// right跳过和val相同的
			continue
		} else {
			// 和value不同的数据滚动向前覆盖
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
