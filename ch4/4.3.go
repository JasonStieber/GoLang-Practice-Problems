package ch4

//import "fmt"

func reverse64BitArray(a *[8]byte) {
	//	fmt.Printf("old: %b ", a)
	for i := 0; i < 4; i++ {
		f := reverseBits(a[i])
		l := reverseBits(a[len(a)-1-i])
		a[i], a[len(a)-1-i] = l, f
		a[len(a)-1-i] = f
	}
	//	fmt.Printf("new:%b\n", a)
}

func reverseBits(b byte) byte {
	newB := byte(0)
	for i := 0; i < 8; i++ {
		newB <<= 1
		if b&1 == 1 {
			newB = newB | byte(1)
		}
		// fmt.Printf("old: %b new: %b \n", b, newB)
		b >>= 1
	}
	return newB
}
