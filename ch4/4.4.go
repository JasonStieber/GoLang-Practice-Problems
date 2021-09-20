package ch4

func singlePassRotate(a *[8]byte, r int) {
	s := *a
	for i := 0; i < 8; i++ {
		s[(i+r)%8] = a[i]
	}
	*a = s
}
