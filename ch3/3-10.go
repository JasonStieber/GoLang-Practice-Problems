package ch3

import (
	"bytes"
)

// 3.10: write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
func Comma(s string) string {
	var buf bytes.Buffer
	for i := range s {
		if (len(s)-i)%3 == 0 && i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		// ok, how do we deal with the commas?
	}
	return buf.String()
}
