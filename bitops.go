package geohashi

// Sieve constants
const (
	s1 = 0x5555555555555555 // 0101010101010101010101010101010101010101010101010101010101010101
	s2 = 0x3333333333333333 // 0011001100110011001100110011001100110011001100110011001100110011
	s3 = 0x0F0F0F0F0F0F0F0F // 0000111100001111000011110000111100001111000011110000111100001111
	s4 = 0x00FF00FF00FF00FF // 0000000011111111000000001111111100000000111111110000000011111111
	s5 = 0x0000FFFF0000FFFF // 0000000000000000111111111111111100000000000000001111111111111111
	s6 = 0x00000000FFFFFFFF // 0000000000000000000000000000000011111111111111111111111111111111
	s7 = 0xAAAAAAAAAAAAAAAA // 1010101010101010101010101010101010101010101010101010101010101010
)

func interleave64(x, y uint64) uint64 {
	x = (x | (x << 16)) & s5
	y = (y | (y << 16)) & s5

	x = (x | (x << 8)) & s4
	y = (y | (y << 8)) & s4

	x = (x | (x << 4)) & s3
	y = (y | (y << 4)) & s3

	x = (x | (x << 2)) & s2
	y = (y | (y << 2)) & s2

	x = (x | (x << 1)) & s1
	y = (y | (y << 1)) & s1

	return x | (y << 1)
}

func deinterleave64(n uint64) (x, y uint64) {
	x = n
	y = n >> 1

	x = (x | (x >> 0)) & s1
	y = (y | (y >> 0)) & s1

	x = (x | (x >> 1)) & s2
	y = (y | (y >> 1)) & s2

	x = (x | (x >> 2)) & s3
	y = (y | (y >> 2)) & s3

	x = (x | (x >> 4)) & s4
	y = (y | (y >> 4)) & s4

	x = (x | (x >> 8)) & s5
	y = (y | (y >> 8)) & s5

	x = (x | (x >> 16)) & s6
	y = (y | (y >> 16)) & s6

	return
}