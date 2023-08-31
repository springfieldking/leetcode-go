package solution

func hammingDistance(x int, y int) int {
	return hammingWeightCleanLastBit(uint32(x ^ y))
}
