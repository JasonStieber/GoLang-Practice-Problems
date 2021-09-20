package ch4

func countBitDifferences(a, b *[32]byte) int {
	dif := 0
	for i := range a {
		dif += countBits(a[i] ^ b[i])
	}
	return dif
}

func countBits(b byte) (count int) {
	for i := 0; i < 8; i++ {
		if b&1 != 0 {
			count++
		}
		b >>= 1
	}
	return count
}
