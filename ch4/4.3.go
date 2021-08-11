package ch4

func reverse64BitArray(a *[8]byte) {
	for i := 0; i < len(a)/2; i++ {
		f := reverseBits(a[i])
		l := reverseBits(a[len(a)-i-1])
		a[i] = l
		a[len(a)-1-i] = f
	}
}

func reverseBits(b byte) byte {
	newB := byte(0)
	for i := 0; i < 8; i++ {
		if b&1 == 1 {
			newB = newB | byte(0000_0001)
		}
		b >>= 1
		newB <<= 1

	}
	return newB
}
