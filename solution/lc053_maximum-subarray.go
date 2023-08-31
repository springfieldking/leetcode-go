package solution

func maxSubArray(nums []int) int {
	// dp arr
	dp := make([]int, len(nums))
	// init
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] < nums[i] {
			// a new sub arr sum
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
