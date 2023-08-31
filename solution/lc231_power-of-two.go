package solution

func isPowerOfTwo(n int) bool {
	// 最高位为1，其他位均为0
	return n != 0 && (n&(n-1) == 0)
}
