package solution

// climbStairs 递归 f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// climbStairsDP
// 1）用DP保存到达此步骤的步骤数量
// 2）状态转换方程 dp[n] = dp[n-1] + dp[n-2]
func climbStairsDP(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	var dp = make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

// climbStairsDPSimplify
// climbStairsDP的简化版本
// 利用p, q配合r组成的动态数组，记录最近三次状态
// 替代原有的dp状态表
func climbStairsDPSimplify(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
