package solution

func twoSum(nums []int, target int) []int {
	tarMap := make(map[int]int)
	for index, num := range nums {
		if peer, ok := tarMap[num]; ok {
			return []int{peer, index}
		} else {
			tarMap[target-num] = index
		}
	}
	return []int{}
}
