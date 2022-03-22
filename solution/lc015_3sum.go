package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var rets [][]int
	var target = 0
	ln := len(nums)
	for i := 0; i < ln; i++ {
		twoSumTar := target - nums[i]
		for j := i + 1; j < ln; j++ {
			oneSumTar := twoSumTar - nums[j]
			for k := j + 1; k < ln; k++ {
				if nums[k] == oneSumTar {
					data := []int{nums[i], nums[j], nums[k]}
					sort.Ints(data)
					has := false
					for _, arr := range rets {
						if arr[0] == data[0] && arr[1] == data[1] && arr[2] == data[2] {
							has = true
							break
						}
					}
					if !has {
						rets = append(rets, data)
					}
				}
			}
		}
	}
	return rets
}
