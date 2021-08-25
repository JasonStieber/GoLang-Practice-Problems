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

import "fmt"

// add is a+b
func add(a, b uint8) uint8 {
	r := uint8(0)
	carry := false
	for i := 0; i < 8; i++ {
		if a>>i&1 != 0 && b>>i&1 != 0 {
			if carry {
				r |= 1 << i
			} else {
				carry = !carry
			}
		} else if a>>i&1 == 0 && b>>i&1 == 0 {
			if carry {
				r |= 1 << i
				carry = !carry
			}
		} else {
			if !carry {
				r |= 1 << i
			}
		}
	}
	fmt.Printf("a:%08b + b:%08b = c:%08b \n", a, b, r)
	return r

}

// mul is a*b
func mul(a, b uint8) uint8 {
	return 0
}

// exp is a to the b
func exp(a, b uint8) uint8 {
	return 0
}

// sub is a-b
func sub(a, b uint8) uint8 {
	return 0
}

// div is a/b
func div(a, b uint8) uint8 {
	return 0
}

// mod is a % b.
func mod(a, b uint8) uint8 {
	return 0
}
