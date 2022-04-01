package solution

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	left := 1
	right := 1
	for ; right < n; right++ {
		if nums[right] == nums[right-1] {
			// right跳过相邻相同的,直到找到相邻不同
			continue
		} else {
			// 用右侧数据覆盖左侧数据
			// 然后left右移一次
			// left <= right
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
