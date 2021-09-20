package ch4

import (
	"testing"
)

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
		{[8]byte{1: 0b1001_0110, 7: 0b1111_0000}, [8]byte{0: 0b0000_1111, 6: 0b0110_1001}},
	} {
		reverse64BitArray(&tt.a)
		if tt.a != tt.want {
			t.Fatalf("Reverse64BitArray should be %b, but is: %b", tt.want, tt.a)
		}
	}
}

func TestSinglePassRotate(t *testing.T) {
	type test = struct {
		a, want [8]byte
		b       int
	}
	for _, tt := range []test{
		{[8]byte{0: 1}, [8]byte{3: 1}, 2},
		{[8]byte{0: 1}, [8]byte{1: 1}, 1},
		{},
		{[8]byte{0: 1, 1: 11, 2: 111}, [8]byte{0: 11, 1: 111, 7: 1}, 6},
		{[8]byte{2: 5, 7: 11}, [8]byte{2: 5, 7: 11}, 0},
	} {
		v := tt.a
		singlePassRotate(&v, tt.b)
		if v != tt.want {
			t.Fatalf("SinglePassRotate(%v, %v) should look like %v, but is: %v", v, tt.b, tt.want, tt.a)
		}
	}
}

// func TestEliminateAdjacentDuplicates(t *testing.T) {
// 	type test = struct {
// 		a, want []string
// 	}
// 	for _, tt := range []test{
// 		{[]string{0: "a", 1: "a"}, []string{0: "a"}},
// 		{[]string{}, []string{}},
// 		{[]string{0: "a", 1: "2", 2: "a", 3: "no"}, []string{0: "a", 1: "2", 2: "a", 3: "no"}},
// 		{[]string{0: "a", 1: "a", 2: "a"}, []string{0: "a"}},
// 		{[]string{0: "", 1: "n", 4: "n"}, []string{0: "", 1: "n", 2: "", 3: "n"}},
// 	} {
// 		got := eliminateAdjacentDuplicates(tt.a)
// 		t.Logf("want: %v, len got: %v", tt.want), len(got))
// 		if len(tt.want) != len(got) {
// 			t.Fatalf("EliminateAdjacentDuplicates(%v) should look like %v, but it is: %v", tt.a, tt.want, got)
// 		}
// 		for i := range got {
// 			if tt.want[i] != got[i] {
// 				t.Fatalf("EliminateAdjacentDuplicates(%v) should look like %v, but it is: %v", tt.a, tt.want, got)
// 			}
// 		}
// 	}
// }

// func TestSquashSpaces(t *testing.T){
// 	type test = struct {
// 		a, want string
// 	}
// 	for _, tt := range []test{
// 		{},
// 		{},
// 		{},
// 		{},
// 	}
// }

func TestReverseCharactersInPlace(t *testing.T) {
	type test = struct {
		input, want string
	}
	for name, tt := range map[string]test{
		"handles unicode": {"🔥--🧊", "🧊--🔥"},
		"trivial":         {},
		"reverses ASCII":  {"foo", "oof"},
	} {
		t.Run(name, func(t *testing.T) {
			v := reverseCharactersInPlace([]byte(tt.input))
			if string(v) != tt.want {
				t.Fatalf("ReverseCharactersInPlace(%v) should look like %v, but is: %c", tt.input, tt.want, v)
			}
		})
	}
}
