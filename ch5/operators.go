// pass the tests in main_test.go
// you can only call functions you define here.
// you can only use the following bitwise operators
//         binary operators:
//             <<, >>, |, &, ^, ==
//      assignment operators:
//            <<=, >>=, |=, &=, ^=
//        prefix unary operator:
//             ^
// I suggest you do them in the order listed.
package operators

// add is a+b
func add(a, b uint8) uint8 {
	r := uint8(0)
	carry := false
	for i := 0; i < 8; i++ {
		bit := uint8(1) << i
		aSet, bSet := a&bit != 0, b&bit != 0

		if aSet && bSet && carry {
			r |= bit
			carry = true
		} else if aSet && bSet || aSet && carry || bSet && carry {
			carry = true
		} else if aSet || bSet || carry {
			carry = false
			r |= bit
		} else {
			carry = false
		}
	}
	return r

}

// mul is a*b
func mul(a, b uint8) uint8 {
	t := uint8(0)
	for i := uint8(0); i != b; i = add(i, 1) {
		t = add(t, a)
	}
	return t
}

// exp is a to the b
func exp(a, b uint8) uint8 {
	t := uint8(1)
	for i := uint8(0); i != b; i = add(i, 1) {
		t = mul(t, a)
	}
	return t
}

// sub is a-b
func sub(a, b uint8) uint8 {
	r := uint8(0)
	borrow := false
	for i := uint8(0); i < 8; i = add(i, 1) {
		aSet, bSet := a&(1<<i) != 0, b&(1<<i) != 0
		if aSet {
			if bSet && borrow {
				r |= 1 << i
				borrow = true
			} else if borrow || bSet {
				borrow = false
			} else {
				r |= 1 << i
				borrow = false
			}
		} else if bSet && borrow {
			borrow = true
		} else if borrow || bSet {
			r |= 1 << i
			borrow = true
		} else {
			borrow = false
		}

	}
	return r
}

func lt(a, b uint8) bool {
	for i := uint8(8); i != 0; sub(i, 1) {
		j := i - 1
		if a&(1<<j) == b&(1<<j) {
			continue
		} else if a&(1<<j) != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func lte(a, b uint8) bool {
	return a == b || lt(a, b)
}

func gt(a, b uint8) bool  { return !lte(a, b) }
func gte(a, b uint8) bool { return !lt(a, b) }

// div is a/b
func div(num, denom uint8) uint8 {
	if denom == 0 {
		panic("You fucked up only Chuck Noris can devide by zero")
	}
	c := uint8(0)
	for ; gt(num, denom); c++ {
		num = sub(num, denom)
	}
	return c
}

// mod is a % b.
func mod(a, b uint8) uint8 {
	return 0
}
