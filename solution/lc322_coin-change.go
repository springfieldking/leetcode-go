package solution

func coinChange(coins []int, amount int) int {
	// init
	max := amount + 1
	dp := make([]int, max)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = max
		// 尝试和上一步的各种组合，去最小值
		for _, coin := range coins {
			if i-coin >= 0 {
				dpi := dp[i-coin] + 1
				if dpi < dp[i] {
					dp[i] = dpi
				}
			}
		}
	}
	if dp[amount] == max {
		return -1
	}
	return dp[amount]
}
