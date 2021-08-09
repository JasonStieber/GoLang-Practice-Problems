package ch3

import (
	"bytes"
	"strings"
)

// 3.11: Enhance comma so that it deals correctly with floating point numbers and an optional sign (+ or -).
func Enhanced_comma(s string) string {

	var byt bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		byt.WriteByte(s[0])
		s = s[1:]
	}

	twoParts := strings.Split(s, ".")
	if len(twoParts) == 2 {
		byt.WriteString(Comma(twoParts[0]))
		byt.WriteByte('.')
		byt.WriteString(twoParts[1])
	} else {
		byt.WriteString(Comma(twoParts[0]))
	}

	return byt.String()
}
