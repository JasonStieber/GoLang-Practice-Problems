package operators

import (
	"fmt"
	"testing"
)

var ops = map[string]struct {
	symbol        string
	op, reference func(uint8, uint8) uint8
}{
	"add": {"+", add, func(a, b uint8) uint8 { return a + b }},
	// "mul": {"*", mul, func(a, b uint8) uint8 { return a * b }},
	// "sub": {"-", sub, func(a, b uint8) uint8 { return a - b }},
	// "div": {"/", div, func(a, b uint8) uint8 { return a / b }},
	// "mod": {"%", mod, func(a, b uint8) uint8 { return a % b }},
	// "pow": {"**", exp, expBySquaring},
}

func Test_Operators(t *testing.T) {
	for name, tt := range ops {
		t.Run(name, func(t *testing.T) {
			for n := 0; n <= 255; n++ {
				for m := 0; m <= 255; m++ {
					exec := func(f func(n, m uint8) uint8) (result uint8, panicked bool) {
						defer func() {
							if p := recover(); p != nil {
								panicked = true
							}
						}()
						result = f(uint8(n), uint8(m))
						return
					}
					got, gotPanicked := exec(tt.op)
					want, wantPanicked := exec(tt.reference)
					format := fmt.Sprintf("%3d %s  %d", n, tt.symbol, m)
					switch {
					case got != want:
						t.Errorf("%s: expected %3d, got %3d", format, want, got)
					case gotPanicked && !wantPanicked:
						t.Errorf("%s: unexpected panic", format)
					case wantPanicked && !gotPanicked:
						t.Errorf("%s: expected a panic, but there was none", format)
					default:

					}
				}
			}

		})
	}
}

// don't bother to write your exponentiation this way; do it in the simplest way you can without worrying about runtime.
func expBySquaring(base, pow uint8) (c uint8) {
	switch {
	case base == 0:
		return 0
	case base == 1, pow == 0:
		return 1
	case base == 1:
		return pow
	default:
		extra := uint8(1)

		for pow > 1 {
			if pow%2 == 0 {
				base, pow = base*base, pow/2
			} else {
				extra, base, pow = extra*extra, base*base, (pow-1)/2
			}
		}
		return base * uint8(extra)
	}
}
