package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		// 相同参数直接跳过，上次已经处理过了
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := 0 - nums[first]
		second := first + 1
		thrid := len(nums) - 1
		for second < thrid {
			// 相同参数直接跳过，上次已经处理过了
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if thrid < len(nums)-1 && nums[thrid] == nums[thrid+1] {
				thrid--
				continue
			}

			// 处理
			if nums[second]+nums[thrid] < target {
				second++
			} else if nums[second]+nums[thrid] > target {
				thrid--
			} else if nums[second]+nums[thrid] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[thrid]})
				second++
				thrid--
			}
		}

	}
	return ans
}
