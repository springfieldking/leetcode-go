package solution

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	left := 1
	right := 1
	for ; right < n; right++ {
		// 相邻相同的跳过,直到找到相邻不同
		// 用右侧数据覆盖左侧数据
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
