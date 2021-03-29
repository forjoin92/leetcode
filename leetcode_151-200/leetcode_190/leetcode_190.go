package leetcode_190

func reverseBits(num uint32) uint32 {
	var res uint32 = num & 1
	for i := 0; i < 31; i++ {
		res <<= 1
		num >>= 1
		res += num & 1
	}
	return res
}
