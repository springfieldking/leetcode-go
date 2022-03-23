package solution

func hammingWeight(num uint32) int {
	ones := 0
	for i := 0; i < 32; i++ {
		if ((1 << i) & num) != 0 {
			ones++
		}
	}
	return ones
}

func hammingWeightCleanLastBit(num uint32) int {
	ones := 0
	for num != 0 {
		num = num & (num - 1)
		ones++
	}
	return ones
}
