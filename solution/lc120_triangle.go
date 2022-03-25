package solution

func minimumTotal(triangle [][]int) int {
	hight := len(triangle)
	dp := make([]int, len(triangle[hight-1]))
	dp[0] = 0
	for i := hight - 1; i >= 0; i-- {
		l := len(triangle[i])
		for j := 0; j < l; j++ {
			if i == hight-1 {
				// 最底层话，最短路径就是自己
				dp[j] = triangle[i][j]
			} else {
				// 否则就是下面分叉元素和当前值的和的最小值
				left := triangle[i][j] + dp[j]
				right := triangle[i][j] + dp[j+1]
				if left < right {
					dp[j] = left
				} else {
					dp[j] = right
				}
			}
		}
	}
	return dp[0]
}
