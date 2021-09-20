package operators ch5

import (
"fmt"
"testing"
)

func Test_ADD(t *testing.T) { testOp(t, ADD) }
func Test_SUB(t *testing.T) { testOp(t, SUB) }
func Test_MUL(t *testing.T) { testOp(t, MUL) }
func Test_DIV(t *testing.T) { testOp(t, DIV) }
func Test_POW(t *testing.T) { testOp(t, POW) }

type OP struct {
	symbol        string
	op, reference func(uint8, uint8) uint8
}

var (
	ADD = OP{"+", add, func(a, b uint8) uint8 { return a + b }}
	SUB = OP{"-", sub, func(a, b uint8) uint8 { return a - b }}
	DIV = OP{"/", div, func(a, b uint8) uint8 { return a / b }}
	MUL = OP{"*", mul, func(a, b uint8) uint8 { return a * b }}
	POW = OP{"**", exp, referenceExp}
)

// for each pair of bytes, compare the result of the operator and the reference operator.
func testOp(t *testing.T, op OP) {
	for n := 0; n <= 4; n++ {
		for m := 0; m <= 4; m++ {
			exec := func(f func(n, m uint8) uint8) (result uint8, panicked bool) {
				defer func() {
					if p := recover(); p != nil {
						panicked = true
					}
				}()
				result = f(uint8(n), uint8(m))
				return
			}
			got, gotPanicked := exec(op.op)
			want, wantPanicked := exec(op.reference)
			format := fmt.Sprintf("%3d %s  %d", n, op.symbol, m)
			switch {
			case got != want:
				t.Errorf("%s: expected %3d, got %3d", format, want, got)
			case gotPanicked && !wantPanicked:
				t.Errorf("%s: unexpected panic", format)
			case wantPanicked && !gotPanicked:
				t.Errorf("%s: expected a panic, but there was none", format)
			default:
				// intentionally blank
			}
		}
	}
	if !t.Failed() {
		t.Logf("%s: OK", t.Name())
	}
}

func referenceExp(base, pow uint8) (c uint8) {
	switch {
	case base == 0 && pow == 0, base == 1:
		return 1
	case base == 0:
		return 0
	default:
		for c = 1; pow > 0; pow-- {
			c *= base
		}
		return c
	}
}
