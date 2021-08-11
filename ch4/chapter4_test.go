package ch4

import "testing"

func TestCountBitDiff(t *testing.T) {
	type test = struct {
		a, b [32]byte
		want int
	}
	for _, tt := range []test{
		{[32]byte{0: 0b1}, [32]byte{0: 0}, 1},
		{},
		{[32]byte{1: 0b111}, [32]byte{3: 0b1011}, 6},
	} {
		if got := countBitDifferences(&tt.a, &tt.b); got != tt.want {
			t.Fatalf("CountBitDifferences(%b, %b) should be %v, but is: %v", tt.a, tt.b, tt.want, got)
		}
	}
}

func TestReverse64BitArray(t *testing.T) {
	type test = struct {
		a, want [8]byte
	}
	for _, tt := range []test{
		{[8]byte{0: 0b1110_0000}, [8]byte{7: 0b0000_0111}},
		{[8]byte{1: 0b1011_0111, 2: 0b1111_0000}, [8]byte{6: 0b1110_1101, 5: 0b0000_1111}},
		{},
		{[8]byte{1: 0b1001_0110, 7: 0b1111_0000}, [8]byte{1: 0b0000_1111, 7: 0b0110_1001}},
	} {
		reverse64BitArray(&tt.a)
		if tt.a != tt.want {
			t.Fatalf("Reverse64BitArray should be %b, but is: %b", tt.want, tt.a)
		}
	}
}

// func TestSinglePassRotate(t *testing.T) {
// 	type test = struct {
// 		a, want, got []byte
// 		b            int
// 	}
// 	for _, tt := range []test{
// 		{},
// 		{},
// 		{},
// 		{},
// 		{},
// 	} {
// 		if got := singlePassRotate(&tt.a, &tt.b); got != tt.want {
// 			t.Fatalf("SinglePassRotate(%b, %b) should look like %v, but is: %v", tt.a, tt.b, tt.want, got)
// 		}
// 	}
// }
